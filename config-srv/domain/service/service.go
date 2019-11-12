package service

import (
	"github.com/micro/go-micro/config/source"
)

type Service interface {
	QueryChangeSet(app, env, cluster string, namespaces ...string) (set *source.ChangeSet, err error)
}

type service struct {
}

var (
	s = &service{}
)

func GetService() Service {
	return s
}

func (s *service) QueryChangeSet(app, env, cluster string, namespaces ...string) (set *source.ChangeSet, err error) {
	return nil, nil
}
