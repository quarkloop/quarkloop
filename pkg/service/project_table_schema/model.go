package table_schema

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetTableListParams struct {
	Context   context.Context
	ProjectId string
}

type GetTableByIdParams struct {
	Context   context.Context
	ProjectId string
	TableId   string
}

type CreateTableParams struct {
	Context   context.Context
	ProjectId string
	Table     *model.Table
}

type CreateBulkTableParams struct {
	Context   context.Context
	ProjectId string
	TableList []model.Table
}

type UpdateTableByIdParams struct {
	Context context.Context
	TableId string
	Table   *model.Table
}

type DeleteTableByIdParams struct {
	Context context.Context
	TableId string
}
