package dto

import "github.com/micro-in-cn/config-server/config-srv/domain/model"

type NSItem struct {
	model.Item
	UpdateType string `json:"updateType,omitempty"`
}

type NS struct {
	AppId     string
	Cluster   string
	Namespace string
}

type NSUpdate struct {
	NS
}
