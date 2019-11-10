package service

import "github.com/micro/go-micro/config/source"

type Service struct {
}



func (s *Service) QueryChangeSet(app, env, cluster string, namespaces ...string) (set *source.ChangeSet, err error) {
	return nil, nil
}
