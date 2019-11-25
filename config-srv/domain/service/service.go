package service

import (
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	"github.com/micro-in-cn/config-server/proto/entry"
	"github.com/micro/go-micro/config/source"
)

var (
	s    *service
	once sync.Once
)

type Service interface {
	QueryChangeSet(app, cluster, namespace string) (set *source.ChangeSet, err error)
	CreateApp(appId, appName string) (int64, error)
	ListApps(appIds ...string) ([]*entry.App, error)
}

type service struct {
	repo     repository.AppRepository
	updateNS chan *watcher.NSUpdate
}

func GetService() Service {
	return s
}

func Init(repository repository.AppRepository, update chan *watcher.NSUpdate) {
	// todo singleton if needs
	once.Do(func() {
		s = &service{
			repo:     repository,
			updateNS: update,
		}
	})
}

func (s *service) QueryChangeSet(app, cluster, namespace string) (set *source.ChangeSet, err error) {
	return nil, nil
}
