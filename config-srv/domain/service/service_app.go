package service

import (
	"github.com/micro-in-cn/config-server/proto/entry"
)

func (s *service) CreateApp(appId, appName string) (int64, error) {
	return 0, nil
}

func (s *service) ListApps(appIds ...string) ([]*entry.App, error) {
	return nil, nil
}
