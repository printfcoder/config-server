package repository

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type Repository interface {
	GlobeRepository
}

type GlobeRepository interface {
	CreateApp(appId, appName string) (int64, error)
	CreateCluster(appId, clusterName string) error
	CreateNamespace(appId, clusterName, namespace string) error
	UpdateItems(appId, clusterName, namespace string, del []*model.Item, update []*model.Item, insert []*model.Item) error
	ListApps(appIds ...string) ([]*model.App, error)
}
