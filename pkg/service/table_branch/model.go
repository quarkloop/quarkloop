package table_branch

import (
	"context"
	"time"
)

type TableBranch struct {
	Id          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Default     bool       `json:"default,omitempty"`
	Type        string     `json:"type,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	CreatedBy   string     `json:"createdBy,omitempty"`
	UpdatedBy   *string    `json:"updatedBy,omitempty"`
}

type GetTableBranchListParams struct {
	Context   context.Context
	ProjectId int
}

type GetTableBranchByIdParams struct {
	Context   context.Context
	ProjectId int
	BranchId  int
}

type CreateTableBranchParams struct {
	Context   context.Context
	ProjectId int
	Branch    *TableBranch
}

type UpdateTableBranchByIdParams struct {
	Context   context.Context
	ProjectId int
	BranchId  int
	Branch    *TableBranch
}

type DeleteTableBranchByIdParams struct {
	Context   context.Context
	ProjectId int
	BranchId  int
}
