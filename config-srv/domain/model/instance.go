package model

import "time"

type Instance struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at" json:"deletedAt"`

	IP          string `gorm:"column:ip" json:"ip"`
	AppID       string `gorm:"column:app_id" json:"appId"`
	ClusterName string `gorm:"column:cluster_name" json:"clusterName"`
}
