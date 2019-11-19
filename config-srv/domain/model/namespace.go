package model

import "time"

type Namespace struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at" json:"deletedAt"`

	Name        string `gorm:"column:name;unique;not null" json:"name"`
	AppID       string `gorm:"column:app_id;unique;not null" json:"appId"`
	ClusterName string `gorm:"column:cluster_name;unique;not null" json:"clusterName"`
}
