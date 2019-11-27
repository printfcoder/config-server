package service

import (
	"github.com/micro-in-cn/config-server/proto/entry"
)

func (s *service) CreateApp(appName string) (*entry.App, error) {
	app, err := s.repo.CreateApp(appName)
	if err != nil {
		return nil, err
	}

	return &entry.App{
		Id:          int64(app.ID),
		Name:        app.Name,
		CreatedTime: app.CreatedAt.Unix(),
		UpdatedTime: app.CreatedAt.Unix(),
	}, nil
}

func (s *service) DeleteApp(appName string) error {
	return s.repo.DeleteApp(appName)
}

func (s *service) ListApps(appNames ...string) ([]*entry.App, error) {
	apps, err := s.repo.ListApps(appNames...)
	if err != nil {
		return nil, err
	}

	var entryApps []*entry.App
	for _, v := range apps {
		entryApps = append(entryApps, &entry.App{
			Id:          int64(v.ID),
			Name:        v.Name,
			CreatedTime: v.CreatedAt.Unix(),
			UpdatedTime: v.CreatedAt.Unix(),
		})
	}

	return entryApps, nil
}
