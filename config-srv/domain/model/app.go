package model

import "time"

type App struct {
	Id          uint64    `json:"id"`
	AppId       string    `json:"appId"`
	AppName     string    `json:"appName"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
