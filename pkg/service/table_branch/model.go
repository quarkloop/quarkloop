package table_branch

import (
	"time"
)

type TableBranch struct {
	Id          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Default     bool   `json:"default,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type GetTableBranchListParams struct {
	ProjectId int32
}

type GetTableBranchByIdParams struct {
	ProjectId int32
	BranchId  int32
}

type CreateTableBranchParams struct {
	ProjectId int32
	Branch    *TableBranch
}

type UpdateTableBranchByIdParams struct {
	ProjectId int32
	BranchId  int32
	Branch    *TableBranch
}

type DeleteTableBranchByIdParams struct {
	ProjectId int32
	BranchId  int32
}
