package table_schema

type Service interface {
	ListTableSchemas(*GetTableSchemaListParams) ([]TableSchema, error)
	GetTableSchemaById(*GetTableSchemaByIdParams) (*TableSchema, error)
	CreateTableSchema(*CreateTableSchemaParams) (*TableSchema, error)
	UpdateTableSchemaById(*UpdateTableSchemaByIdParams) error
	DeleteTableSchemaById(*DeleteTableSchemaByIdParams) error
}
