package model

import "time"

type File struct {
	AppId      string    `json:"appId,omitempty"`
	InstanceId string    `json:"appInstanceId,omitempty"`
	Id         string    `json:"id,omitempty"`
	Enable     *bool     `json:"enable,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
