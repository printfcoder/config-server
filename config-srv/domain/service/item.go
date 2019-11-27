package service

import (
	"fmt"
	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	"sync"
	"time"

	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	cjson "github.com/micro-in-cn/config-server/config-srv/util/json"
	"github.com/micro-in-cn/config-server/proto/entry"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/util/log"
)

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
		err = fmt.Errorf("[UpdateNSItems] update items error: %s", err.Error())
		log.Error(err)
		return
	}

	dtoNSUpdate := &dto.NSUpdate{
		NS: dto.NS{
			AppId:     app,
			Cluster:   cluster,
			Namespace: namespace,
		},
	}
	var wg sync.WaitGroup
	wg.Add(2)

	// pub event to sibling nodes
	go func() {
		// some biz logic here
		s.updateNSForPub <- dtoNSUpdate
		wg.Done()
	}()

	// tell Watcher sendMsg back to the clients
	go func() {
		set, err := s.QueryChangeSet(app, cluster, namespace)
		if err != nil {
			err = fmt.Errorf("[UpdateNSItems] query the newest set error: %s", err.Error())
			log.Error(err)
			return
		}

		s.updateNSForWatcher <- &watcher.NSUpdate{
			NSUpdate:        dtoNSUpdate,
			NewestDataBytes: set.Data,
		}
		wg.Done()
	}()

	// todo timeout
	wg.Wait()

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
