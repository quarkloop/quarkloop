package model

import (
	"encoding/json"
	"time"
)

type App struct {
	AppId      string          `json:"appId"`
	InstanceId string          `json:"instanceId"`
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Type       int             `json:"type,omitempty"`
	Icon       string          `json:"icon,omitempty"`
	Metadata   json.RawMessage `json:"metadata,omitempty"`
	CreatedAt  time.Time       `json:"createdAt,omitempty"`
	UpdatedAt  time.Time       `json:"updatedAt,omitempty"`
}
