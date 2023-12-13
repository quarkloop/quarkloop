package table_schema

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetTableSchemaListParams struct {
	Context   context.Context
	ProjectId string
}

type GetTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId string
	SchemaId  string
}

type CreateTableSchemaParams struct {
	Context   context.Context
	ProjectId string
	Schema    *model.TableSchema
}

type CreateBulkTableSchemaParams struct {
	Context    context.Context
	ProjectId  string
	SchemaList []model.TableSchema
}

type UpdateTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId string
	SchemaId  string
	Schema    *model.TableSchema
}

type DeleteTableSchemaByIdParams struct {
	Context   context.Context
	ProjectId string
	SchemaId  string
}
