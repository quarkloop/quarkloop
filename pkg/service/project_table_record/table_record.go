package table_record

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	ListTableRecords(*GetTableListParams) ([]model.TableWithRelationCount, error)
	GetTableRecordById(*GetTableByIdParams) (*model.TableWithRelationCount, error)
	CreateTable(*CreateTableParams) (*model.Table, error)
	CreateBulkTable(*CreateBulkTableParams) (int64, error)
	UpdateTableById(*UpdateTableByIdParams) error
	DeleteTableById(*DeleteTableByIdParams) error
}
