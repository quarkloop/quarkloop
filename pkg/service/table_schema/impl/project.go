package table_schema_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema/store"
)

type tableSchemaService struct {
	store store.TableSchemaStore
}

func NewTableSchemaService(ds store.TableSchemaStore) table_schema.Service {
	return &tableSchemaService{
		store: ds,
	}
}

func (s *tableSchemaService) ListTableSchemas(ctx context.Context, p *table_schema.GetTableSchemaListParams) ([]table_schema.TableSchema, error) {
	schemaList, err := s.store.ListTableSchemas(ctx, p.ProjectId)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

func (s *tableSchemaService) GetTableSchemaById(ctx context.Context, p *table_schema.GetTableSchemaByIdParams) (*table_schema.TableSchema, error) {
	schema, err := s.store.GetTableSchemaById(ctx, p.ProjectId, p.SchemaId)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func (s *tableSchemaService) CreateTableSchema(ctx context.Context, p *table_schema.CreateTableSchemaParams) (*table_schema.TableSchema, error) {
	schema, err := s.store.CreateTableSchema(ctx, p.ProjectId, p.Schema)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func (s *tableSchemaService) UpdateTableSchemaById(ctx context.Context, p *table_schema.UpdateTableSchemaByIdParams) error {
	err := s.store.UpdateTableSchemaById(ctx, p.ProjectId, p.SchemaId, p.Schema)
	return err
}

func (s *tableSchemaService) DeleteTableSchemaById(ctx context.Context, p *table_schema.DeleteTableSchemaByIdParams) error {
	err := s.store.DeleteTableSchemaById(ctx, p.ProjectId, p.SchemaId)
	return err
}
