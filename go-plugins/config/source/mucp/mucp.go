package mucp

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/config/source"
)

var (
	DefaultPath        = "/micro/config"
	DefaultServiceName = "go.micro.config"
)

type mucpSource struct {
	serviceName string
	path        string
	opts        source.Options
	client      client.Client
	service     proto.SourceService
}

func (m *mucpSource) Read() (set *source.ChangeSet, err error) {
	req := m.client.NewRequest(m.serviceName, "Source.Read", &proto.ReadRequest{Path: m.path})
	if err != nil {
		return nil, err
	}

	out := new(proto.ReadResponse)
	err = m.client.Call(context.Background(), req, out)
	if err != nil {
		return nil, err
	}

	return toChangeSet(out.ChangeSet), nil
}

func (m *mucpSource) Watch() (w source.Watcher, err error) {
	stream, err := m.service.Watch(context.Background(), &proto.WatchRequest{Path: m.path})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	return newWatcher(stream)
}

func (m *mucpSource) String() string {
	return "proto"
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	addr := DefaultServiceName
	path := DefaultPath

	if options.Context != nil {
		a, ok := options.Context.Value(serviceNameKey{}).(string)
		if ok {
			addr = a
		}
		p, ok := options.Context.Value(pathKey{}).(string)
		if ok {
			path = p
		}
	}

	s := &mucpSource{
		serviceName: addr,
		path:        path,
		opts:        options,
		client:      *cmd.DefaultOptions().Client,
	}

	s.service = proto.NewSourceService("go.micro.config", s.client)

	s.client.Init()

	return s
}
