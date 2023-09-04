package types

import "time"

type File struct {
	Enable    bool      `json:"enable"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
