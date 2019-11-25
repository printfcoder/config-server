package watcher

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-micro/config/source"
)

var (
	mu    sync.Mutex
	w     *watcher
	chMap = make(map[string][]chan *source.ChangeSet)
)

type Watcher interface {
	Watch(app, cluster string, namespace string) (ch chan *source.ChangeSet)
}

type watcher struct {
	updateChan chan *NSUpdate
}

func (w *watcher) run() {
	f := func(ch chan *source.ChangeSet, cs *NSUpdate) {
		set := &source.ChangeSet{
			Data:      cs.NewestBytes,
			Format:    "json",
			Source:    "mucp", // or mysql or some names else
			Timestamp: time.Now(),
		}
		set.Checksum = set.Sum()
		ch <- set
	}

	for cs := range w.updateChan {
		csKey := getChKey(cs.AppId, cs.Cluster, cs.Namespace)
		if chs, ok := chMap[csKey]; ok {
			for _, ch := range chs {
				go f(ch, cs)
			}
		}
	}
}

func (w *watcher) Watch(app, cluster string, namespace string) (ch chan *source.ChangeSet) {
	// for locating which chan should be pushed event into
	key := getChKey(app, cluster, namespace)

	mu.Lock()
	ch = make(chan *source.ChangeSet)
	chMap[key] = append(chMap[key], ch)
	mu.Unlock()

	return
}

func GetWatcher() Watcher {
	return w
}

func NewWatcher(opts ...Option) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	w = &watcher{
		updateChan: options.UpdateChan,
	}
	go w.run()
	return
}

func getChKey(app, cluster, namespace string) string {
	return string(md5.New().Sum([]byte(fmt.Sprintf("%s%s%s", app, cluster, namespace))))
}
