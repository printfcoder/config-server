package mysql

import (
	"context"

	"github.com/micro/go-micro/config/source"
)

type appName struct{}
type clusterName struct{}
type namespace struct{}
type serviceRef struct{}

type QueryService interface {
	QueryChangeSet(app, cluster, namespace string) (set *source.ChangeSet, err error)
	Watch(app, cluster string, namespaces string) (ch chan *source.ChangeSet)
}

func WithApp(app string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, appName{}, app)
	}
}

func WithClusterName(cluster string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, clusterName{}, cluster)
	}
}

func WithNamespace(ns ...string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, namespace{}, ns)
	}
}

func WithQueryService(service QueryService) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, serviceRef{}, service)
	}
}
