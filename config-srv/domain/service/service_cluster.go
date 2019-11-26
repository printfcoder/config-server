package service

import "github.com/micro-in-cn/config-server/proto/entry"

func (s *service) CreateCluster(appName, clusterName string) (*entry.Cluster, error) {
	cluster, err := s.repo.CreateCluster(appName, clusterName)
	if err != nil {
		return nil, err
	}

	return &entry.Cluster{
		Id:          int64(cluster.ID),
		Name:        cluster.Name,
		AppName:     cluster.AppName,
		CreatedTime: cluster.CreatedAt.Unix(),
		UpdatedTime: cluster.UpdatedAt.Unix(),
	}, nil
}

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
