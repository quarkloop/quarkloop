package table_schema

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	ListTableSchemas(*GetTableListParams) ([]model.TableWithRelationCount, error)
	GetTableSchemaById(*GetTableByIdParams) (*model.TableWithRelationCount, error)
	CreateTableSchema(*CreateTableParams) (*model.Table, error)
	UpdateTableSchemaById(*UpdateTableByIdParams) error
	DeleteTableSchemaById(*DeleteTableByIdParams) error
}
