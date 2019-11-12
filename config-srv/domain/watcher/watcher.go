package watcher

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/micro/go-micro/config/source"
)

var (
	w     = &watcher{}
	mu    sync.Mutex
	chMap = make(map[string][]*source.ChangeSet)
	nsMap = make(map[string][]*source.ChangeSet)
)

func Init() {

}

type Watcher interface {
	Watch(app, env, cluster string, namespaces ...string) (ch chan *source.ChangeSet)
}

type watcher struct {
}

func (w *watcher) Watch(app, env, cluster string, namespaces ...string) (ch chan *source.ChangeSet) {
	// for locating which chans should be pushed event into
	key := getChKey(app, env, cluster, namespaces...)
	// for listening the changes of configs in every namespace
	nsKeys := getNSKeys(app, cluster, namespaces...)

	mu.Lock()
	if sets, ok := chMap[key]; !ok {
		chMap[key] = []chan *source.ChangeSet{make(chan *source.ChangeSet)}
	}
	mu.Unlock()

	return nil
}

func GetWatcher() Watcher {
	return w
}

func getChKey(app, env, cluster string, namespaces ...string) string {
	sort.Strings(namespaces)
	return fmt.Sprintf("%s%s%s%s", app, env, cluster, strings.Join(namespaces, ""))
}

func getNSKeys(app, cluster string, namespaces ...string) []string {
	ret := make([]string, 0, len(namespaces))
	for _, v := range namespaces {
		ret = append(ret, fmt.Sprintf("%s%s%s", app, cluster, v))
	}

	return ret
}
