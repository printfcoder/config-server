package watcher

import (
	"crypto/md5"
	"fmt"
	"sync"

	"github.com/micro/go-micro/config/source"
)

var (
	w     = &watcher{}
	mu    sync.Mutex
	chMap = make(map[string][]chan *source.ChangeSet)
	nsMap = make(map[string][]chan *source.ChangeSet)
)

func Init() {

}

type Watcher interface {
	Watch(app, env, cluster string, namespaces string) (ch chan *source.ChangeSet)
}

type watcher struct {
}

func (w *watcher) Run() {

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

func GetWatcher() Watcher {
	return w
}

func getChKey(app, env, cluster, namespace string) string {
	return string(md5.New().Sum([]byte(fmt.Sprintf("%s%s%s%s", app, env, cluster, namespace))))
}
