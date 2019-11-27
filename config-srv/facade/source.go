package facade

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
)

type source struct{}

var (
	_            proto.SourceHandler = &source{}
	watchExitMap                     = map[string]chan bool{}
	mu           sync.RWMutex
)

func (s *source) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) (err error) {
	ns, err := parsePath(req.Path)
	if err != nil {
		log.Errorf("[Read] parse path err: %s", err.Error())
		return
	}

	set, err := se.QueryChangeSet(ns.AppId, ns.Cluster, ns.Namespace)
	if err != nil {
		log.Errorf("[Read] QueryChangeSet err: %s", err.Error())
		return
	}

	rsp.ChangeSet = &proto.ChangeSet{
		Data:      set.Data,
		Checksum:  set.Checksum,
		Format:    "json",
		Source:    "mucp",
		Timestamp: time.Now().Unix()}
	return
}

func (s *source) Watch(ctx context.Context, req *proto.WatchRequest, server proto.Source_WatchStream) (err error) {
	ns, err := parsePath(req.Path)
	if err != nil {
		log.Errorf("[Watch] parse path err: %s", err.Error())
		return
	}

	w := watcher.GetWatcher()
	ch := w.Watch(ns.AppId, ns.Cluster, ns.Namespace)

	// shall we support holding multi connects in the remote service?
	// todo check the mt below is existed
	mt, _ := metadata.FromContext(ctx)
	exit := make(chan bool)
	go func(remote string) {
		mu.RLock()
		defer mu.RUnlock()

		// delete the old one and let Watch return
		if oldOne, ok := watchExitMap[remote]; ok {
			delete(watchExitMap, remote)
			oldOne <- true
		}

		// add the new one
		mu.Lock()
		watchExitMap[remote] = exit
		mu.Unlock()
	}(mt["Remote"])

	for {
		select {
		case set := <-ch:
			{
				rsp := &proto.WatchResponse{
					ChangeSet: &proto.ChangeSet{
						Data:      set.Data,
						Checksum:  set.Checksum,
						Format:    "json",
						Source:    "mucp",
						Timestamp: time.Now().Unix()},
				}
				if err = server.SendMsg(rsp); err != nil {
					log.Logf("[Watch] watch files error，%s", err)
					return err
				}
			}
		case <-exit:
			// do not close this socket! just let GC recycle the current goroutine
			// the client will reuse it for the next request, if we close it now, the client will never get any message
			// until recreating a new stream connection.
			return
		}
	}

	return
}

func parsePath(path string) (ns *dto.NS, err error) {
	// protocol format: "VERSION1:#{APP}/#{CLUSTER}/#{Namespace}"
	ns = &dto.NS{}
	paths := strings.Split(path, ":")
	if len(paths) != 2 {
		return nil, fmt.Errorf("[parsePath] path doesn't contain any version. the standard path is same as \"VERSION1:#{APP}/#{CLUSTER}/#{Namespace}\"")
	}

	// ignore the version now, we dont need it to do anything at present. is just a placeholder
	paths = strings.Split(paths[1], "/")
	if len(paths) != 3 {
		return nil, fmt.Errorf("[parsePath] path is not standard. the standard path is same as \"VERSION1:#{APP}/#{CLUSTER}/#{Namespace}\"")
	}

	ns.AppId = paths[0]
	ns.Cluster = paths[1]
	ns.Namespace = paths[2]

	return
}
