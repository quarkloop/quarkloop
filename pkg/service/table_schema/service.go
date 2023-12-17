package table_schema

import "context"

type Service interface {
	ListTableSchemas(context.Context, *GetTableSchemaListParams) ([]TableSchema, error)
	GetTableSchemaById(context.Context, *GetTableSchemaByIdParams) (*TableSchema, error)
	CreateTableSchema(context.Context, *CreateTableSchemaParams) (*TableSchema, error)
	UpdateTableSchemaById(context.Context, *UpdateTableSchemaByIdParams) error
	DeleteTableSchemaById(context.Context, *DeleteTableSchemaByIdParams) error
}
