package table_branch

import (
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
	ProjectId int
}

type GetTableBranchByIdParams struct {
	ProjectId int
	BranchId  int
}

type CreateTableBranchParams struct {
	ProjectId int
	Branch    *TableBranch
}

type UpdateTableBranchByIdParams struct {
	ProjectId int
	BranchId  int
	Branch    *TableBranch
}

type DeleteTableBranchByIdParams struct {
	ProjectId int
	BranchId  int
}
