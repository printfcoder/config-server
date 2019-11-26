package facade

import (
	"context"
	"fmt"

	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	proto "github.com/micro-in-cn/config-server/proto/entry"
	"github.com/prometheus/common/log"
)

type entry struct{}

var (
	_ proto.EntryHandler = entry{}
)

func (e entry) CreateApp(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) (err error) {
	app := req.GetApp()
	if app != nil {
		newApp, err := service.GetService().CreateApp(app.GetName())
		if err != nil {
			err = fmt.Errorf("[CreateApp] create app error: %s", err.Error())
			goto Error
		}
		rsp.App.Id = newApp.Id
		return nil
	} else {
		err = fmt.Errorf("[CreateApp] req.app is nil")
		goto Error
	}

Error:
	{
		log.Warn(err)
		rsp.Error = &proto.Error{
			// todo，define united error code
			Msg: err.Error(),
		}
		return nil
	}
}

func (e entry) ListApps(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) GetApp(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateCluster(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) GetCluster(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) ListClusters(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) PullInstances(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateNamespace(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) UpdateItems(context.Context, *proto.EntryRequest, *proto.EntryResponse) error {
	panic("implement me")
}
