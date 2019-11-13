package mysql

import (
	"context"
	"github.com/micro/go-micro/config/source"
)

type appName struct{}
type envName struct{}
type clusterName struct{}
type namespace struct{}

func WithApp(app string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, appName{}, app)
	}
}

func WithEnv(env string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, envName{}, env)
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
