package model

import "time"

type Cluster struct {
	ID        uint       `gorm:"column:id;          primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at"               json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at"               json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at"   json:"deletedAt"`

	AppName string `gorm:"column:app_name; not null; unique_index:cluster_uindex" json:"appName"`
	Name    string `gorm:"column:name;     not null; unique_index:cluster_uindex" json:"name"`
}
