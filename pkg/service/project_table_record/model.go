package table_record

import (
	"context"
)

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
