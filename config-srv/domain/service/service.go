package service

import "github.com/micro/go-micro/config/source"

type Service struct {
}

func (s *Service) QueryChangeSet(app, env, cluster string, namespaces ...string) (set *source.ChangeSet, err error) {
	return nil, nil
}

func (s *Service) Watch(app, env, cluster string, namespace string, ch chan *source.ChangeSet) (err error) {
	return nil
}
