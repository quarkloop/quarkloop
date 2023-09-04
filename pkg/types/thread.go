package types

import "time"

type Thread struct {
	Type      ThreadType `json:"thread_type"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type ThreadType int

const (
	Normal ThreadType = 1
	Inline ThreadType = 2
)
