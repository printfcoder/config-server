package repository

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type Repository interface {
	GlobalRepository
}

type GlobalRepository interface {
	CreateApp(appName string) (*model.App, error)
	CreateCluster(appName, clusterName string) (*model.Cluster, error)
	CreateNamespace(appName, clusterName, namespaceName string) (*model.Namespace, error)
	UpdateItems(appName, clusterName, namespace string, del []*model.Item, update []*model.Item, insert []*model.Item) error
	ListApps(appNames ...string) ([]*model.App, error)
}
