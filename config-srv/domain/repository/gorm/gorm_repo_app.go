package gorm

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type appRepo struct{}

func (a appRepo) CreateApp(appId, appName string) (int64, error) {
	db.Select("SELECT * FROM")
	panic("implement me")
}

func (a appRepo) ListApps(appIds ...string) ([]*model.App, error) {
	panic("implement me")
}

func (a appRepo) CreateEnv(appId, envName string) error {
	panic("implement me")
}

func (a appRepo) CreateCluster(appId, envName, clusterName string) error {
	panic("implement me")
}

func (a appRepo) CreateNamespace(appId, envName, clusterName, namespace string) error {
	panic("implement me")
}

func (a appRepo) CreateItems(appId, envName, clusterName, namespace string, items map[string]string) error {
	panic("implement me")
}

func (a appRepo) DeleteApp(appId, appName string) (int64, error) {
	panic("implement me")
}

func (a appRepo) DeleteEnv(appId, envName string) error {
	panic("implement me")
}

func (a appRepo) DeleteCluster(appId, envName, clusterName string) error {
	panic("implement me")
}

func (a appRepo) DeleteNamespace(appId, envName, clusterName, namespace string) error {
	panic("implement me")
}

func (a appRepo) DeleteItems(appId, envName, clusterName, namespace string, keys ...string) error {
	panic("implement me")
}

func (a appRepo) UpdateItems(appId, envName, clusterName, namespace string, items map[string]string) error {
	panic("implement me")
}
