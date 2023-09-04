package types

import "time"

type App struct {
	Name      string    `json:"name"`
	Type      int       `json:"type"`
	Status    AppStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Icon      string    `json:"icon"`
}

type AppStatus int

const (
	On       AppStatus = 0
	Off      AppStatus = 1
	Archived AppStatus = 2
)
