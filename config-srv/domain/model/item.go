package model

import "time"

type Item struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at" json:"deletedAt"`

	Key         string `gorm:"column:key;not null" json:"key"`
	Value       string `gorm:"column:value;not null" json:"value"`
	NamespaceID int    `gorm:"column:namespace_id;not null" json:"namespaceId"`
}
