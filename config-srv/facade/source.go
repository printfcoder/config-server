package facade

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-in-cn/config-server/config-srv/loader"
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/util/log"
)

type source struct{}

func (s source) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) (err error) {
	appName := loader.ParsePath(req.Path)
	switch appName {
	case "micro", "extra":
		rsp.ChangeSet = loader.GetConfig(appName)
		return
	default:
		err = fmt.Errorf("[Read] the first path is invalid")
		return
	}
}

func (s source) Watch(ctx context.Context, req *proto.WatchRequest, server proto.Source_WatchStream) (err error) {
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
