package mucp

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/source"
)

type watcher struct {
	stream client.Stream
}

func newWatcher(stream client.Stream) (*watcher, error) {
	return &watcher{stream: stream}, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	rsp, err := w.stream.Recv()
	if err != nil {
		return nil, err
	}
	return toChangeSet(rsp.ChangeSet), nil
}

func (w *watcher) Stop() error {
	return w.stream.Close()
}
