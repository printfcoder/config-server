package gorm

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type appRepo struct {
}

func (a appRepo) CreateApp(appId, appName string) (int64, error) {
	db.Select("SELECT * FROM")
	panic("implement me")
}

func (a appRepo) ListApps(appIds ...string) ([]*model.App, error) {
	panic("implement me")
}
