package model

import "time"

type Item struct {
	Id          uint64    `json:"id"`
	NamespaceId uint64    `json:"namespaceId"`
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
