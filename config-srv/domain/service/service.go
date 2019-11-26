package service

import (
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/domain/model"
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
	CreateApp(appName string) (int64, error)
	ListApps(appNames ...string) ([]*entry.App, error)
}

type service struct {
	repo     repository.GlobalRepository
	updateNS chan *watcher.NSUpdate
}

func GetService() Service {
	return s
}

func Init(repository repository.GlobalRepository, update chan *watcher.NSUpdate) {
	// todo singleton if needs
	once.Do(func() {
		s = &service{
			repo:     repository,
			updateNS: update,
		}
	})
}

func (s *service) QueryChangeSet(app, cluster, namespace string) (set *source.ChangeSet, err error) {
	return
}

func (s *service) UpdateNSItems(app, cluster, namespace string, items []*entry.Item) (err error) {
	if len(items) == 0 {
		return
	}

	del, update, insert := groupItemsForUpdate(items)
	// todo log
	return s.repo.UpdateItems(app, cluster, namespace, del, update, insert)
}

func groupItemsForUpdate(items []*entry.Item) (del []*model.Item, update []*model.Item, insert []*model.Item) {
	for _, item := range items {
		// mode one
		m := convertEntryItemToModel(item)
		// TODO item.UpdateType
		switch 1 {
		case 1:
			del = append(del, m)
		case 2:
			update = append(update, m)
		case 3:
			insert = append(insert, m)
		}
	}

	return
}

func convertEntryItemToModel(item *entry.Item) (ret *model.Item) {
	return &model.Item{}
}
