package service

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
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

var _ Service = &service{}

type Service interface {
	QueryChangeSet(app, cluster, namespace string) (set *source.ChangeSet, err error)
	UpdateNSItems(app, cluster, namespace string, items []*entry.Item) (err error)

	CreateApp(appName string) (*entry.App, error)
	DeleteApp(appName string) error
	ListApps(appNames ...string) ([]*entry.App, error)

	CreateCluster(appName, clusterName string) (*entry.Cluster, error)
	DeleteCluster(appName, clusterName string) error
	ListClusters(appName string) ([]*entry.Cluster, error)

	CreateNamespace(appName, clusterName, namespaceName string) (*entry.Namespace, error)
	DeleteNamespace(appName, clusterName, namespaceName string) error
	ListNamespace(appName, clusterName string) ([]*entry.Namespace, error)
}

type service struct {
	repo               repository.GlobalRepository
	updateNSForWatcher chan<- *watcher.NSUpdate
	updateNSForPub     chan<- *dto.NSUpdate
}

func GetService() Service {
	return s
}

func Init(repository repository.GlobalRepository, updateNS chan<- *watcher.NSUpdate, updateNSForPub chan<- *dto.NSUpdate) {
	// todo singleton if needs
	once.Do(func() {
		s = &service{
			repo:               repository,
			updateNSForWatcher: updateNS,
			updateNSForPub:     updateNSForPub,
		}
	})
}
