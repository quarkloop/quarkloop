package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

type TableSchemaStore interface {
	ListTableSchemas(ctx context.Context, projectId int) ([]table_schema.TableSchema, error)
	GetTableSchemaById(ctx context.Context, projectId int, schemaId string) (*table_schema.TableSchema, error)
	CreateTableSchema(ctx context.Context, projectId int, schema *table_schema.TableSchema) (*table_schema.TableSchema, error)
	UpdateTableSchemaById(ctx context.Context, projectId int, schemaId string, schema *table_schema.TableSchema) error
	DeleteTableSchemaById(ctx context.Context, projectId int, schemaId string) error
}

type tableSchemaStore struct {
	Conn *pgx.Conn
}

func NewTableSchemaStore(conn *pgx.Conn) *tableSchemaStore {
	return &tableSchemaStore{
		Conn: conn,
	}
}
