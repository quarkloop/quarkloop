package project_table_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table_record "github.com/quarkloop/quarkloop/pkg/service/project_table_record"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableService struct {
	dataStore *repository.Repository
}

func NewTableService(ds *repository.Repository) table_record.Service {
	return &tableService{
		dataStore: ds,
	}
}

func (s *tableService) ListTableRecords(p *table_record.GetTableListParams) ([]model.TableWithRelationCount, error) {
	projectList, err := s.dataStore.ListTableRecords(p.Context, p.ProjectId)
	return projectList, err
}

func (s *tableService) GetTableRecordById(p *table_record.GetTableByIdParams) (*model.TableWithRelationCount, error) {
	project, err := s.dataStore.GetTableRecordById(p.Context, p.ProjectId, p.TableId)
	return project, err
}

func (s *tableService) CreateTable(p *table_record.CreateTableParams) (*model.Table, error) {
	project, err := s.dataStore.CreateTable(p.Context, p.ProjectId, p.Table)
	return project, err
}

func (s *tableService) CreateBulkTable(p *table_record.CreateBulkTableParams) (int64, error) {
	project, err := s.dataStore.CreateBulkTable(p.Context, p.ProjectId, p.TableList)
	return project, err
}

func (s *tableService) UpdateTableById(p *table_record.UpdateTableByIdParams) error {
	err := s.dataStore.UpdateTableById(p.Context, p.TableId, p.Table)
	return err
}

func (s *tableService) DeleteTableById(p *table_record.DeleteTableByIdParams) error {
	err := s.dataStore.DeleteTableById(p.Context, p.TableId)
	return err
}
