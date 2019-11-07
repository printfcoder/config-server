package facade

import (
	"context"
	"github.com/micro-in-cn/config-server/proto/config"
)

type manage struct {
}

func (c manage) CreateApp(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) ListApps(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) GetApp(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) CreateCluster(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) GetCluster(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) ListClusters(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) CreateEnv(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) GetEnv(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) ListEnvs(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) CreateInstance(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) CreateNamespace(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) CreateItem(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) UpdateItem(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c manage) DeleteItems(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}
