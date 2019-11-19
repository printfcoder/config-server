package repository

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type Repository interface {
	AppRepository
}

type AppRepository interface {
	CreateApp(appId, appName string) (int64, error)
	ListApps(appIds ...string) ([]*model.App, error)
}
