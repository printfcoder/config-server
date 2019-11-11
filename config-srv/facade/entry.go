package facade

import (
	"context"
	"fmt"

	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	proto "github.com/micro-in-cn/config-server/proto/entry"
)

type entry struct {
}

var (
	appService = &service.App{}
)

func (e entry) CreateApp(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) (err error) {
	app := req.GetApp()
	if app != nil {
		id, err := appService.CreateApp(app.GetAppId(), app.GetAppName())
		if err != nil {
			err = fmt.Errorf("[CreateApp] create app error: %s", err)
			return
		}
		rsp.App.Id = id
		return
	}

	return
}

func (e entry) ListApps(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) GetApp(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateCluster(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) GetCluster(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) ListClusters(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateEnv(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) GetEnv(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) ListEnvs(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) PullInstances(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateNamespace(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) CreateItem(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) UpdateItem(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}

func (e entry) DeleteItems(ctx context.Context, req *proto.EntryRequest, rsp *proto.EntryResponse) error {
	panic("implement me")
}
