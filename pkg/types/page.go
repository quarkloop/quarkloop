package types

import "time"

type Page struct {
	Name       string    `json:"name"`
	Entrypoint bool      `json:"entrypoint"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
