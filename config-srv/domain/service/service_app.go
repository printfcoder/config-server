package service

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/proto/entry"
)

type App struct {
	repo repository.AppRepository
}

func (s *App) CreateApp(appId, appName string) (int64, error) {
	return 0, nil
}

func (s *App) ListApps(appIds ...string) ([]*entry.App, error) {
	return nil, nil
}

func NewAppService(repository repository.AppRepository) *App {
	// todo singleton if needs
	return &App{repo: repository}
}
