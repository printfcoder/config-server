package model

import "time"

type Item struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at" json:"deletedAt"`

	Key         string `gorm:"column:key" json:"key"`
	Value       string `gorm:"column:value" json:"value"`
	NamespaceID int    `gorm:"column:namespace_id" json:"namespaceId"`
}
