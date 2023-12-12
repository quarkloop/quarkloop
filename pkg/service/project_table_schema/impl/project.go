package project_table_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table_schema "github.com/quarkloop/quarkloop/pkg/service/project_table_schema"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableService struct {
	dataStore *repository.Repository
}

func NewTableService(ds *repository.Repository) table_schema.Service {
	return &tableService{
		dataStore: ds,
	}
}

func (s *tableService) ListTableSchemas(p *table_schema.GetTableListParams) ([]model.TableWithRelationCount, error) {
	projectList, err := s.dataStore.ListTableSchemas(p.Context, p.ProjectId)
	return projectList, err
}

func (s *tableService) GetTableSchemaById(p *table_schema.GetTableByIdParams) (*model.TableWithRelationCount, error) {
	project, err := s.dataStore.GetTableSchemaById(p.Context, p.ProjectId, p.TableId)
	return project, err
}

func (s *tableService) CreateTableSchema(p *table_schema.CreateTableParams) (*model.Table, error) {
	project, err := s.dataStore.CreateTableSchema(p.Context, p.ProjectId, p.Table)
	return project, err
}

func (s *tableService) UpdateTableSchemaById(p *table_schema.UpdateTableByIdParams) error {
	err := s.dataStore.UpdateTableSchemaById(p.Context, p.TableId, p.Table)
	return err
}

func (s *tableService) DeleteTableSchemaById(p *table_schema.DeleteTableByIdParams) error {
	err := s.dataStore.DeleteTableSchemaById(p.Context, p.TableId)
	return err
}
