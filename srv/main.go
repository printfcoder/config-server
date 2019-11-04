package main

import (
	"context"
	"fmt"
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro-in-cn/config-server/srv/config"
	"github.com/micro-in-cn/config-server/srv/db"
	"github.com/micro-in-cn/config-server/srv/loader"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"sync"
)

var (
	mux        sync.RWMutex
	configMaps = make(map[string]*proto.ChangeSet)
)

type Source struct{}

func main() {
	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.config"),
	)

	// 注册服务
	proto.RegisterSourceHandler(service.Server(), new(Source))

	service.Init(micro.Action(func(ctx *cli.Context) {
		config.Init()
		db.Init()
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
