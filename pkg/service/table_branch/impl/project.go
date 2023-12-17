package table_branch_impl

import (
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch/store"
)

type tableBranchService struct {
	store store.TableBranchStore
}

func NewTableBranchService(ds store.TableBranchStore) table_branch.Service {
	return &tableBranchService{
		store: ds,
	}
}

func (s *tableBranchService) ListTableBranches(p *table_branch.GetTableBranchListParams) ([]table_branch.TableBranch, error) {
	branchList, err := s.store.ListTableBranches(p.Context, p.ProjectId)
	if err != nil {
		return nil, err
	}

	return branchList, nil
}

func (s *tableBranchService) GetTableBranchById(p *table_branch.GetTableBranchByIdParams) (*table_branch.TableBranch, error) {
	branch, err := s.store.GetTableBranchById(p.Context, p.ProjectId, p.BranchId)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (s *tableBranchService) CreateTableBranch(p *table_branch.CreateTableBranchParams) (*table_branch.TableBranch, error) {
	branch, err := s.store.CreateTableBranch(p.Context, p.ProjectId, p.Branch)
	if err != nil {
		return nil, err
	}

	return branch, nil
}

func (s *tableBranchService) UpdateTableBranchById(p *table_branch.UpdateTableBranchByIdParams) error {
	err := s.store.UpdateTableBranchById(p.Context, p.ProjectId, p.BranchId, p.Branch)
	return err
}

func (s *tableBranchService) DeleteTableBranchById(p *table_branch.DeleteTableBranchByIdParams) error {
	err := s.store.DeleteTableBranchById(p.Context, p.ProjectId, p.BranchId)
	return err
}