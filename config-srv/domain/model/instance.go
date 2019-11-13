package model

import "time"

type Instance struct {
	Id          uint64    `json:"id"`
	AppId       string    `json:"appId"`
	ClusterName string    `json:"clusterName"`
	Ip          string    `json:"ip"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
