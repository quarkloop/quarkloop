package table_branch_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	table_branch "github.com/quarkloop/quarkloop/pkg/service/project_table_branch"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableBranchService struct {
	dataStore *repository.Repository
}

func NewTableBranchService(ds *repository.Repository) table_branch.Service {
	return &tableBranchService{
		dataStore: ds,
	}
}

func (s *tableBranchService) ListTableBranches(p *table_branch.GetTableBranchListParams) ([]model.TableBranch, error) {
	branchList, err := s.dataStore.ListTableBranches(p.Context, p.ProjectId)
	return branchList, err
}

func (s *tableBranchService) GetTableBranchById(p *table_branch.GetTableBranchByIdParams) (*model.TableBranch, error) {
	branch, err := s.dataStore.GetTableBranchById(p.Context, p.ProjectId, p.BranchId)
	return branch, err
}

func (s *tableBranchService) CreateTableBranch(p *table_branch.CreateTableBranchParams) (*model.TableBranch, error) {
	branch, err := s.dataStore.CreateTableBranch(p.Context, p.ProjectId, p.Branch)
	return branch, err
}

func (s *tableBranchService) UpdateTableBranchById(p *table_branch.UpdateTableBranchByIdParams) error {
	err := s.dataStore.UpdateTableBranchById(p.Context, p.ProjectId, p.BranchId, p.Branch)
	return err
}

func (s *tableBranchService) DeleteTableBranchById(p *table_branch.DeleteTableBranchByIdParams) error {
	err := s.dataStore.DeleteTableBranchById(p.Context, p.ProjectId, p.BranchId)
	return err
}
