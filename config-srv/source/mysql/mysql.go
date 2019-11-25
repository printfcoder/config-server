package mysql

import (
	"github.com/micro/go-micro/config/source"
)

var (
	DefaultCluster = "default"
)

type mysqlSource struct {
	app, cluster string
	namespace    string
	opts         source.Options
	client       QueryService
}

func (m *mysqlSource) Read() (*source.ChangeSet, error) {
	return m.client.QueryChangeSet(m.app, m.cluster, m.namespace)
}

func (m *mysqlSource) Watch() (source.Watcher, error) {
	// confirm and prepare data is here
	cs, err := m.Read()
	if err != nil {
		return nil, err
	}

	return newWatcher(m.client, cs, m.app, m.cluster, m.namespace)
}

func (m *mysqlSource) String() string {
	return "mysql"
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	cluster := DefaultCluster
	app, ns := "", ""
	var queryService QueryService
	ok := false

	if options.Context != nil {
		app, ok = options.Context.Value(appName{}).(string)
		if !ok {
			panic("App name is necessary for loading configurations! plz check on the Platform-Web which name it is.")
		}

		ns, ok = options.Context.Value(namespace{}).(string)
		if !ok {
			panic("Namespace is necessary for loading configurations! it decides which domain configs you need.")
		}

		ns, ok = options.Context.Value(namespace{}).(string)
		if !ok {
			panic("Namespace is necessary for loading configurations! it decides which domain configs you need.")
		}

		queryService, ok = options.Context.Value(serviceRef{}).(QueryService)
		if !ok {
			panic("QueryService is necessary for loading configurations! it decides where to query the config data.")
		}
	} else {
		panic("Mysql source config error!")
	}

	s := &mysqlSource{
		app:       app,
		cluster:   cluster,
		namespace: ns,
		opts:      options,
		client:    queryService,
	}

	return s
}
