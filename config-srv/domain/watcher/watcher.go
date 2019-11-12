package watcher

import "github.com/micro/go-micro/config/source"

var (
	w = &watcher{}
)

func Init() {

}

type Watcher interface {
	Watch(app, env, cluster string, namespaces ...string) (ch chan *source.ChangeSet)
}

type watcher struct {
}

func (w *watcher) Watch(app, env, cluster string, namespaces ...string) (ch chan *source.ChangeSet) {
	return nil
}

func GetWatcher() Watcher {
	return w
}
