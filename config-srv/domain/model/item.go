package model

import "time"

type Item struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Key         string    `gorm:"column:key" json:"key"`
	Value       string    `gorm:"column:value" json:"value"`
	NamespaceID int       `gorm:"column:namespace_id" json:"namespaceId"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updatedTime"`
	Deleted     int       `gorm:"column:deleted" json:"deleted"`
}
