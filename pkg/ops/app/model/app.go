package model

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Type      int             `json:"type,omitempty"`
	Status    AppStatus       `json:"status,omitempty"`
	Icon      string          `json:"icon,omitempty"`
	Metadata  json.RawMessage `json:"metadata,omitempty"`
	CreatedAt time.Time       `json:"createdAt,omitempty"`
	UpdatedAt time.Time       `json:"updatedAt,omitempty"`
}

type AppStatus int

const (
	On       AppStatus = 0
	Off      AppStatus = 1
	Archived AppStatus = 2
)
