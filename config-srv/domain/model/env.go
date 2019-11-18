package model

import "time"

type Env struct {
	ID          int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name        string    `gorm:"column:name" json:"name"`
	AppID       string    `gorm:"column:app_id" json:"appId"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updatedTime"`
	Deleted     int       `gorm:"column:deleted" json:"deleted"`
}
