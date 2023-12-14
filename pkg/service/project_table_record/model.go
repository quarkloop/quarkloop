package table_record

import (
	"context"
)

type GetTableRecordListParams struct {
	Context   context.Context
	TableType string
	ProjectId string
	BranchId  string
}

type GetTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId string
	BranchId  string
	RecordId  string
}

type CreateTableRecordParams struct {
	Context   context.Context
	TableType string
	ProjectId string
	BranchId  string
	Record    interface{}
}

type UpdateTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId string
	BranchId  string
	RecordId  string
	Record    interface{}
}

type DeleteTableRecordByIdParams struct {
	Context   context.Context
	TableType string
	ProjectId string
	BranchId  string
	RecordId  string
}
