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

func (s *service) DeleteCluster(appName, clusterName string) error {
	return s.repo.DeleteCluster(appName, clusterName)
}

func (s *service) ListClusters(appName string) ([]*entry.Cluster, error) {
	clusters, err := s.repo.ListClusters(appName)
	if err != nil {
		return nil, err
	}

	var entryClusters []*entry.Cluster
	for _, v := range clusters {
		entryClusters = append(entryClusters, &entry.Cluster{
			Id:          int64(v.ID),
			Name:        v.Name,
			AppName:     v.AppName,
			CreatedTime: v.CreatedAt.Unix(),
			UpdatedTime: v.UpdatedAt.Unix(),
		})
	}

	return entryClusters, nil
}
