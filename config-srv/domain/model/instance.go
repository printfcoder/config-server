package model

import "time"

type Instance struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	IP          string    `gorm:"column:ip" json:"ip"`
	AppID       string    `gorm:"column:app_id" json:"appId"`
	ClusterName string    `gorm:"column:cluster_name" json:"clusterName"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updatedTime"`
}
