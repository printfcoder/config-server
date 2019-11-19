package service

import (
	"github.com/micro-in-cn/config-server/proto/entry"
)

type App struct {
}

func (s *App) CreateApp(appId, appName string) (int64, error) {
	return 0, nil
}

func (s *App) ListApps(appIds ...string) ([]*entry.App, error) {
	return nil, nil
}
