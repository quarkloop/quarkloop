package model

import "time"

type AppInstance struct {
	ID        string    `json:"id,omitempty"`
	AppId     string    `json:"projectId"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
