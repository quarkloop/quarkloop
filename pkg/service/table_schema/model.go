package table_schema

import (
	"context"
	"encoding/json"
	"time"
)

type TableSchema struct {
	Id          int             `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Type        int             `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`
	CreatedAt   time.Time       `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty"`
	CreatedBy   string          `json:"createdBy,omitempty"`
	UpdatedBy   *string         `json:"updatedBy,omitempty"`
}

type GetTableSchemaListParams struct {
	Context   context.Context
	ProjectId int
}

type GetTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId int
	SchemaId  string
}

type CreateTableSchemaParams struct {
	Context   context.Context
	ProjectId int
	Schema    *TableSchema
}

type CreateBulkTableSchemaParams struct {
	Context    context.Context
	ProjectId  string
	SchemaList []TableSchema
}

type UpdateTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId int
	SchemaId  string
	Schema    *TableSchema
}

type DeleteTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId int
	SchemaId  string
}
