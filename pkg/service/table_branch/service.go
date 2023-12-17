package table_branch

import "context"

type Service interface {
	ListTableBranches(context.Context, *GetTableBranchListParams) ([]TableBranch, error)
	GetTableBranchById(context.Context, *GetTableBranchByIdParams) (*TableBranch, error)
	CreateTableBranch(context.Context, *CreateTableBranchParams) (*TableBranch, error)
	UpdateTableBranchById(context.Context, *UpdateTableBranchByIdParams) error
	DeleteTableBranchById(context.Context, *DeleteTableBranchByIdParams) error
}
