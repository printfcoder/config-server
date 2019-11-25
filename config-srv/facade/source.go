package facade

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/util/log"
)

type source struct{}

var _ proto.SourceHandler = &source{}

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

	set, err := se.QueryChangeSet(ns.AppId, ns.Cluster, ns.Namespace)
	if err != nil {
		log.Errorf("[Watch] QueryChangeSet err: %s", err.Error())
		return
	}

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
