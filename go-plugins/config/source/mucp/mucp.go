package mucp

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/source"
)

type mucpSource struct {
	serviceName string
	path        string
	opts        source.Options
	client      client.Client
}

func (m *mucpSource) Read(set *source.ChangeSet, err error) {

}

func (g *mucpSource) Watch() (source.Watcher, error) {

}

func (m *mucpSource) String() string {
	return "mucp"
}
