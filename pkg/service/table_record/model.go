package table_record

import (
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

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type DocumentRecord struct {
	Id          int             `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Type        int             `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type FormRecord struct {
	Id          int             `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Type        int             `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type PaymentRecord struct {
	Id          int             `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Type        int             `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
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
	TableType string
	ProjectId int
	BranchId  int
}

type GetTableRecordByIdParams struct {
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
}

type CreateTableRecordParams struct {
	TableType string
	ProjectId int
	BranchId  int
	Record    interface{}
}

type UpdateTableRecordByIdParams struct {
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
	Record    interface{}
}

type DeleteTableRecordByIdParams struct {
	TableType string
	ProjectId int
	BranchId  int
	RecordId  string
}
