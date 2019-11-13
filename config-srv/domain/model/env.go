package model

import "time"

type Env struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	AppId       string    `json:"appId"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
