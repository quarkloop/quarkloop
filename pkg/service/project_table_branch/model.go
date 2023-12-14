package table_branch

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetTableBranchListParams struct {
	Context   context.Context
	ProjectId string
}

type GetTableBranchByIdParams struct {
	Context   context.Context
	ProjectId string
	BranchId  int
}

type CreateTableBranchParams struct {
	Context   context.Context
	ProjectId string
	Branch    *model.TableBranch
}

type UpdateTableBranchByIdParams struct {
	Context   context.Context
	ProjectId string
	BranchId  int
	Branch    *model.TableBranch
}

type DeleteTableBranchByIdParams struct {
	Context   context.Context
	ProjectId string
	BranchId  int
}
