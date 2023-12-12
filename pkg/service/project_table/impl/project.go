package project_table_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table "github.com/quarkloop/quarkloop/pkg/service/project_table"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableService struct {
	dataStore *repository.Repository
}

func NewTableService(ds *repository.Repository) table.Service {
	return &tableService{
		dataStore: ds,
	}
}

func (s *tableService) ListTableRecords(p *table.GetTableListParams) ([]model.TableWithRelationCount, error) {
	projectList, err := s.dataStore.ListTableRecords(p.Context, p.ProjectId)
	return projectList, err
}

func (s *tableService) GetTableRecordById(p *table.GetTableByIdParams) (*model.TableWithRelationCount, error) {
	project, err := s.dataStore.GetTableRecordById(p.Context, p.ProjectId, p.TableId)
	return project, err
}

func (s *tableService) CreateTable(p *table.CreateTableParams) (*model.Table, error) {
	project, err := s.dataStore.CreateTable(p.Context, p.ProjectId, p.Table)
	return project, err
}

func (s *tableService) CreateBulkTable(p *table.CreateBulkTableParams) (int64, error) {
	project, err := s.dataStore.CreateBulkTable(p.Context, p.ProjectId, p.TableList)
	return project, err
}

func (s *tableService) UpdateTableById(p *table.UpdateTableByIdParams) error {
	err := s.dataStore.UpdateTableById(p.Context, p.TableId, p.Table)
	return err
}

func (s *tableService) DeleteTableById(p *table.DeleteTableByIdParams) error {
	err := s.dataStore.DeleteTableById(p.Context, p.TableId)
	return err
}
