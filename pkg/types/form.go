package types

import "time"

type Form struct {
	Name       string      `json:"name"`
	Fields     []FormField `json:"fields"`
	FieldCount int         `json:"field_count"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type FormField struct{}
