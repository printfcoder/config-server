package model

import (
	"time"
)

type App struct {
	ID        uint       `gorm:"column:id;          primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at"               json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at"               json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at"   json:"deletedAt"`

	Name string `gorm:"column:name; not null; unique_index:app_uindex" json:"name"`
}
