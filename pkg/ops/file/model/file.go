package model

import "time"

type File struct {
	Id        string    `json:"id,omitempty"`
	Enable    *bool     `json:"enable,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
