package model

import "time"

type Item struct {
	ID        uint       `gorm:"column:id;          primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at"               json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at"               json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at"   json:"deletedAt"`

	AppName       string `gorm:"column:app_name;       not null; unique_index:item_uindex" json:"appName"`
	ClusterName   string `gorm:"column:cluster_name;   not null; unique_index:item_uindex" json:"clusterName"`
	NamespaceName string `gorm:"column:namespace_name; not null; unique_index:item_uindex" json:"namespaceName"`
	Key           string `gorm:"column:key;            not null; unique_index:item_uindex" json:"key"`
	Value         string `gorm:"column:value;"                                             json:"value"`
}
