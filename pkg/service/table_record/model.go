package table_record

import (
	"context"
	"encoding/json"
	"time"
)

type MainRecord struct {
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

type DocumentRecord struct {
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

type FormRecord struct {
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

type PaymentRecord struct {
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

type MainRecordWithRelationCount struct {
	Table     MainRecord `json:"table"`
	Relations struct {
		Document int `json:"document,omitempty"`
		Form     int `json:"form,omitempty"`
		Payment  int `json:"payment,omitempty"`
	}
}

type GetTableRecordListParams struct {
	Context   context.Context
	TableType string
	ProjectId int
	BranchId  int
}

type GetTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
}

type CreateTableRecordParams struct {
	Context   context.Context
	TableType string
	ProjectId int
	BranchId  int
	Record    interface{}
}

type UpdateTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
	Record    interface{}
}

type DeleteTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
}
