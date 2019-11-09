package mysql

import (
	"github.com/micro/go-micro/config/source"
)

type watcher struct {
	ch   chan *source.ChangeSet
	exit chan bool
}

func newWatcher(cw client.Watcher) (source.Watcher, error) {
	return &watcher{cw: cw}, nil
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
