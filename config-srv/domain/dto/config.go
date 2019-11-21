package dto

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type NamespaceItems struct {
	model.Item
}

type NSUpdate struct {
	AppId     string
	Cluster   string
	Namespace string
	Env       string
}
