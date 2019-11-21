package mysql

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	cwc "github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	"github.com/micro/go-micro/config/source"
)

var (
	DefaultCluster = "default"
)

type mysqlSource struct {
	app, env, cluster string
	namespace         string
	opts              source.Options
	client            service.Service
	wc                cwc.Watcher
}

func (m *mysqlSource) Read() (*source.ChangeSet, error) {
	return m.client.QueryChangeSet(m.app, m.env, m.cluster, m.namespace)
}

func (m *mysqlSource) Watch() (source.Watcher, error) {
	// confirm and prepare data is here
	cs, err := m.Read()
	if err != nil {
		return nil, err
	}

	return newWatcher(m.wc, cs, m.app, m.env, m.cluster, m.namespace)
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
	app, env, ns := "", "", ""
	ok := false

	if options.Context != nil {
		app, ok = options.Context.Value(appName{}).(string)
		if !ok {
			panic("App name is necessary for loading configurations! plz check on the Platform-Web which name it is.")
		}

		env, ok = options.Context.Value(envName{}).(string)
		if !ok {
			panic("Env is necessary for loading configurations! eg. It probably should be DEV or FAT or UAT...")
		}

		ns, ok = options.Context.Value(namespace{}).(string)
		if !ok {
			panic("Namespace is necessary for loading configurations! it decides which domain configs you need.")
		}
	} else {
		panic("Mysql source config error!")
	}

	s := &mysqlSource{
		app:       app,
		env:       env,
		cluster:   cluster,
		namespace: ns,
		opts:      options,
		client:    service.GetService(),
		// TODO
		//wc:        cwc.GetWatcher(),
	}

	return s
}
