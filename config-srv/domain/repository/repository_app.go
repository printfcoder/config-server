package repository

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type Repository interface {
	AppRepository
}

type AppRepository interface {
	CreateApp(appId, appName string) (int64, error)
	CreateEnv(appId, envName string) error
	CreateCluster(appId, envName, clusterName string) error
	CreateNamespace(appId, envName, clusterName, namespace string) error
	CreateItems(appId, envName, clusterName, namespace string, items map[string]string) error

	DeleteApp(appId, appName string) (int64, error)
	DeleteEnv(appId, envName string) error
	DeleteCluster(appId, envName, clusterName string) error
	DeleteNamespace(appId, envName, clusterName, namespace string) error
	DeleteItems(appId, envName, clusterName, namespace string, keys ...string) error

	UpdateItems(appId, envName, clusterName, namespace string, items map[string]string) error

	ListApps(appIds ...string) ([]*model.App, error)
}
