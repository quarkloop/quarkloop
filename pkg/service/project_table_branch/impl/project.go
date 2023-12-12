package project_table_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table_branch "github.com/quarkloop/quarkloop/pkg/service/project_table_branch"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableService struct {
	dataStore *repository.Repository
}

func NewTableService(ds *repository.Repository) table_branch.Service {
	return &tableService{
		dataStore: ds,
	}
}

func (s *tableService) ListTableBranches(p *table_branch.GetTableListParams) ([]model.TableWithRelationCount, error) {
	projectList, err := s.dataStore.ListTableBranches(p.Context, p.ProjectId)
	return projectList, err
}

func (s *tableService) GetTableBranchById(p *table_branch.GetTableByIdParams) (*model.TableWithRelationCount, error) {
	project, err := s.dataStore.GetTableBranchById(p.Context, p.ProjectId, p.TableId)
	return project, err
}

func (s *tableService) CreateTableBranch(p *table_branch.CreateTableParams) (*model.Table, error) {
	project, err := s.dataStore.CreateTableBranch(p.Context, p.ProjectId, p.Table)
	return project, err
}

func (s *tableService) UpdateTableBranchById(p *table_branch.UpdateTableByIdParams) error {
	err := s.dataStore.UpdateTableBranchById(p.Context, p.TableId, p.Table)
	return err
}

func (s *tableService) DeleteTableBranchById(p *table_branch.DeleteTableByIdParams) error {
	err := s.dataStore.DeleteTableBranchById(p.Context, p.TableId)
	return err
}
