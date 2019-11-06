package main

import (
	"context"
	"fmt"
	"github.com/micro-in-cn/config-server/srv/broker"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro-in-cn/config-server/srv/config"
	"github.com/micro-in-cn/config-server/srv/db"
	"github.com/micro-in-cn/config-server/srv/loader"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

var (
	mux        sync.RWMutex
	configMaps = make(map[string]*proto.ChangeSet)
)

type Source struct{}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.config"),
		micro.WrapCall(),
	)

	proto.RegisterSourceHandler(service.Server(), new(Source))

	service.Init(micro.Action(func(ctx *cli.Context) {
		config.Init()
		db.Init()
		broker.Init()
	}))

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func (s Source) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) (err error) {
	appName := loader.ParsePath(req.Path)
	switch appName {
	case "micro", "extra":
		rsp.ChangeSet = loader.GetConfig(appName)
		return
	default:
		err = fmt.Errorf("[Read] the first path is invalid")
		return
	}

	return
}

func (s Source) Watch(ctx context.Context, req *proto.WatchRequest, server proto.Source_WatchStream) (err error) {
	appName := loader.ParsePath(req.Path)
	rsp := &proto.WatchResponse{
		ChangeSet: loader.GetConfig(appName),
	}
	if err = server.SendMsg(rsp); err != nil {
		log.Logf("[Watch] watch files error，%s", err)
		return err
	}
	return
}
