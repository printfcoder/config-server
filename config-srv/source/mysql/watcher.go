package mysql

import (
	"sync"

	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/util/log"
)

type watcher struct {
	sync.RWMutex
	ch   chan *source.ChangeSet
	cs   *source.ChangeSet
	wc   QueryService //nolint
	exit chan bool
}

func newWatcher(wc QueryService, cs *source.ChangeSet, app, cluster string, namespace string) (source.Watcher, error) {
	w := &watcher{
		cs:   cs,
		ch:   make(chan *source.ChangeSet),
		exit: make(chan bool),
	}

	ch := wc.Watch(app, cluster, namespace)

	go w.run(wc, ch)

	return w, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case cs := <-w.ch:
		return cs, nil
	case <-w.exit:
		return nil, source.ErrWatcherStopped
	}
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	default:
		close(w.exit)
	}
	return nil
}

func (w *watcher) handle(cs *source.ChangeSet) {
	cs.Checksum = cs.Sum()

	w.Lock()
	w.cs = cs
	w.Unlock()

	w.ch <- cs
}

func (w *watcher) run(wc QueryService, ch chan *source.ChangeSet) {
	for {
		select {
		case rsp := <-ch:
			w.handle(rsp)
		case <-w.exit:
			if err := w.Stop(); err != nil {
				log.Errorf("[Watcher] stop error: %s", err)
			}
			return
		}
	}
}
