package mucp

import (
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/config/source"
)

type watcher struct {
	stream proto.Source_WatchService
}

func newWatcher(stream proto.Source_WatchService) (*watcher, error) {
	return &watcher{stream: stream}, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	var rsp proto.WatchResponse
	err := w.stream.RecvMsg(&rsp)
	if err != nil {
		return nil, err
	}
	return toChangeSet(rsp.ChangeSet), nil
}

func (w *watcher) Stop() error {
	return w.stream.Close()
}
