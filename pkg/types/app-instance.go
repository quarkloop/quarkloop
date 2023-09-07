package types

import "time"

type AppInstance struct {
	Id        string    `json:"id"`
	AppId     string    `json:"appId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
