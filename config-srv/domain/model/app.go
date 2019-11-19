package model

import "time"

type App struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	AppID       string    `gorm:"column:app_id" json:"appId"`
	AppName     string    `gorm:"column:app_name" json:"appName"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updatedTime"`
	Deleted     int       `gorm:"column:deleted" json:"deleted"`
}
