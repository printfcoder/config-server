package service

import "github.com/micro-in-cn/config-server/proto/entry"

func (s *service) CreateNamespace(appName, clusterName, namespaceName string) (*entry.Namespace, error) {
	namespace, err := s.repo.CreateNamespace(appName, clusterName, namespaceName)
	if err != nil {
		return nil, err
	}

	return &entry.Namespace{
		Id:          int64(namespace.ID),
		AppName:     namespace.AppName,
		Name:        namespace.Name,
		ClusterName: namespace.ClusterName,
		CreatedTime: namespace.CreatedAt.Unix(),
		UpdatedTime: namespace.UpdatedAt.Unix(),
	}, nil
}

func (s *service) DeleteNamespace(appName, clusterName, namespaceName string) error {
	return s.repo.DeleteNamespace(appName, clusterName, namespaceName)
}

func (s *service) ListNamespace(appName, clusterName string) ([]*entry.Namespace, error) {
	namespaces, err := s.repo.ListNamespace(appName, clusterName)
	if err != nil {
		return nil, err
	}

	var entryNamespaces []*entry.Namespace
	for _, v := range namespaces {
		entryNamespaces = append(entryNamespaces, &entry.Namespace{
			Id:          int64(v.ID),
			AppName:     v.AppName,
			Name:        v.Name,
			ClusterName: v.ClusterName,
			CreatedTime: v.CreatedAt.Unix(),
			UpdatedTime: v.UpdatedAt.Unix(),
		})
	}

	return entryNamespaces, nil
}
