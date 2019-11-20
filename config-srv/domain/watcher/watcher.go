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
	chMap = make(map[string][]chan *source.ChangeSet)
)

type Watcher interface {
	Watch(app, env, cluster string, namespaces string) (ch chan *source.ChangeSet)
}

type watcher struct {
	updateChan chan *NSUpdate
}

func (w *watcher) run() {
	select {
	case cs := <-w.updateChan:
		{
			csKey := getChKey(cs.AppId, cs.Env, cs.Cluster, cs.Namespace)
			if chs, ok := chMap[csKey]; ok {
				f := func(ch chan *source.ChangeSet) {
					set := &source.ChangeSet{
						Data:      cs.NewestBytes,
						Format:    "json",
						Source:    "mucp", // or mysql or some names else
						Timestamp: time.Now(),
					}
					set.Checksum = set.Sum()
					ch <- set
				}

				for _, ch := range chs {
					go f(ch)
				}
			}
		}
	}
}

func (w *watcher) Watch(app, env, cluster string, namespace string) (ch chan *source.ChangeSet) {
	// for locating which chan should be pushed event into
	key := getChKey(app, env, cluster, namespace)

	mu.Lock()
	ch = make(chan *source.ChangeSet)
	chMap[key] = append(chMap[key], ch)
	mu.Unlock()

	return
}

func NewWatcher(opts Options) Watcher {
	w := &watcher{
		updateChan: opts.UpdateChan,
	}
	go w.run()
	return w
}

func getChKey(app, env, cluster, namespace string) string {
	return string(md5.New().Sum([]byte(fmt.Sprintf("%s%s%s%s", app, env, cluster, namespace))))
}
