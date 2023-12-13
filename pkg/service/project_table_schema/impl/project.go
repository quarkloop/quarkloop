package table_schema_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table_schema "github.com/quarkloop/quarkloop/pkg/service/project_table_schema"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableSchemaService struct {
	dataStore *repository.Repository
}

func NewTableSchemaService(ds *repository.Repository) table_schema.Service {
	return &tableSchemaService{
		dataStore: ds,
	}
}

func (s *tableSchemaService) ListTableSchemas(p *table_schema.GetTableSchemaListParams) ([]model.TableSchema, error) {
	schemaList, err := s.dataStore.ListTableSchemas(p.Context, p.ProjectId)
	return schemaList, err
}

func (s *tableSchemaService) GetTableSchemaById(p *table_schema.GetTableSchemaByIdParams) (*model.TableSchema, error) {
	schema, err := s.dataStore.GetTableSchemaById(p.Context, p.SchemaId)
	return schema, err
}

func (s *tableSchemaService) CreateTableSchema(p *table_schema.CreateTableSchemaParams) (*model.TableSchema, error) {
	schema, err := s.dataStore.CreateTableSchema(p.Context, p.ProjectId, p.Schema)
	return schema, err
}

func (s *tableSchemaService) UpdateTableSchemaById(p *table_schema.UpdateTableSchemaByIdParams) error {
	err := s.dataStore.UpdateTableSchemaById(p.Context, p.ProjectId, p.SchemaId, p.Schema)
	return err
}

func (s *tableSchemaService) DeleteTableSchemaById(p *table_schema.DeleteTableSchemaByIdParams) error {
	err := s.dataStore.DeleteTableSchemaById(p.Context, p.ProjectId, p.SchemaId)
	return err
}
