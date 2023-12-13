package table_schema

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	ListTableSchemas(*GetTableSchemaListParams) ([]model.TableSchema, error)
	GetTableSchemaById(*GetTableSchemaByIdParams) (*model.TableSchema, error)
	CreateTableSchema(*CreateTableSchemaParams) (*model.TableSchema, error)
	UpdateTableSchemaById(*UpdateTableSchemaByIdParams) error
	DeleteTableSchemaById(*DeleteTableSchemaByIdParams) error
}
