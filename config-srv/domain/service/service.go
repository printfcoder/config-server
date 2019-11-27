package service

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"sync"
	"time"

	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	cjson "github.com/micro-in-cn/config-server/config-srv/util/json"
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

func (s *service) QueryChangeSet(appName, cluster, namespace string) (set *source.ChangeSet, err error) {
	items, err := s.repo.ListItems(appName, cluster, namespace)
	if err != nil {
		err = fmt.Errorf("[QueryChangeSet] select items error: %s. ", err.Error())
		log.Error(err)
		return
	}

	// convert to json
	jsonLists := make([]string, len(items))
	for _, item := range items {
		jsonMap := map[string]interface{}{}
		jsonLists = append(jsonLists, cjson.DotStringToJSON(item.Key, item.Value, jsonMap))
	}

	// todo combine jsonLists into a json string
	jsonStr := cjson.MergeJSONs(jsonLists)

	set = &source.ChangeSet{
		Data:      []byte(jsonStr),
		Format:    "json",
		Source:    "mucp",
		Timestamp: time.Now(),
	}

	// todo supports sign?
	set.Checksum = set.Sum()

	return
}

func (s *service) UpdateNSItems(app, cluster, namespace string, items []*entry.Item) (err error) {
	if len(items) == 0 {
		return
	}

	del, update, insert := groupItemsForUpdate(items)
	err = s.repo.UpdateItems(app, cluster, namespace, del, update, insert)
	if err != nil {
		log.Error("[UpdateNSItems] update items error: %s", err.Error())
		return
	}

	return
}

func groupItemsForUpdate(items []*entry.Item) (del []*model.Item, update []*model.Item, insert []*model.Item) {
	for _, item := range items {
		// mode one
		m := convertEntryItemToModel(item)
		switch item.UpdateType {
		case entry.UpdateType_DELETE:
			del = append(del, m)
		case entry.UpdateType_UPDATE:
			update = append(update, m)
		case entry.UpdateType_INSERT:
			insert = append(insert, m)
		}
	}

	return
}

func convertEntryItemToModel(item *entry.Item) (ret *model.Item) {
	return &model.Item{
		Key:   item.Key,
		Value: item.Value,
	}
}
