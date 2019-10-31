package mucp

import (
	"context"
	mucp "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/source"
)

var (
	DefaultPath        = "/micro/config"
	DefaultServiceName = "go.micro.configService"
)

type mucpSource struct {
	serviceName string
	path        string
	opts        source.Options
	client      client.Client
}

func (m *mucpSource) Read() (set *source.ChangeSet, err error) {
	req := m.client.NewRequest(m.serviceName, "", &mucp.ReadRequest{Path: m.path})
	if err != nil {
		return nil, err
	}

	out := new(mucp.ReadResponse)
	err = m.client.Call(context.Background(), req, out)
	if err != nil {
		return nil, err
	}

	return toChangeSet(out.ChangeSet), nil
}

func (g *mucpSource) Watch() (w source.Watcher, err error) {
	return nil, nil
}

func (m *mucpSource) String() string {
	return "mucp"
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

	return &mucpSource{
		serviceName: addr,
		path:        path,
		opts:        options,
	}
}
