package model

import "time"

type Namespace struct {
	Id          uint64    `json:"id"`
	AppId       string    `json:"app_id"`
	Name        string    `json:"name"`
	ClusterName uint64    `json:"clusterName"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
