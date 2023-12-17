package table_record

import "context"

type Service interface {
	ListTableRecords(context.Context, *GetTableRecordListParams) (interface{}, error)
	GetTableRecordById(context.Context, *GetTableRecordByIdParams) (interface{}, error)
	CreateTableRecord(context.Context, *CreateTableRecordParams) (interface{}, error)
	UpdateTableRecordById(context.Context, *UpdateTableRecordByIdParams) error
	DeleteTableRecordById(context.Context, *DeleteTableRecordByIdParams) error
}
