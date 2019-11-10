package facade

import (
	"context"

	"github.com/micro-in-cn/config-server/proto/config"
)

type entry struct {
}

func (c entry) CreateApp(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) ListApps(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) GetApp(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) CreateCluster(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) GetCluster(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) ListClusters(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) CreateEnv(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) GetEnv(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) ListEnvs(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) PullInstances(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) CreateNamespace(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) CreateItem(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) UpdateItem(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}

func (c entry) DeleteItems(context.Context, *config.ConfigRequest, *config.ConfigResponse) error {
	panic("implement me")
}
