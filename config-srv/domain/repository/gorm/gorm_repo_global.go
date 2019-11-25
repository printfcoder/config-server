package gorm

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type globalRepo struct{}

func (g globalRepo) CreateApp(appId, appName string) (int64, error) {
	panic("implement me")
}

func (g globalRepo) CreateCluster(appId, clusterName string) error {
	panic("implement me")
}

func (g globalRepo) CreateNamespace(appId, clusterName, namespace string) error {
	panic("implement me")
}

func (g globalRepo) UpdateItems(appId, clusterName, namespace string, del []*model.Item, update []*model.Item, insert []*model.Item) error {
	return nil
}

func (g globalRepo) ListApps(appIds ...string) ([]*model.App, error) {
	panic("implement me")
}
