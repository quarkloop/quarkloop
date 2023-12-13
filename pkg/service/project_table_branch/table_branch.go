package table_branch

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	ListTableBranches(*GetTableBranchListParams) ([]model.TableBranch, error)
	GetTableBranchById(*GetTableBranchByIdParams) (*model.TableBranch, error)
	CreateTableBranch(*CreateTableBranchParams) (*model.TableBranch, error)
	UpdateTableBranchById(*UpdateTableBranchByIdParams) error
	DeleteTableBranchById(*DeleteTableBranchByIdParams) error
}
