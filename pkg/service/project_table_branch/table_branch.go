package table_branch

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	ListTableBranches(*GetTableListParams) ([]model.TableWithRelationCount, error)
	GetTableBranchById(*GetTableByIdParams) (*model.TableWithRelationCount, error)
	CreateTableBranch(*CreateTableParams) (*model.Table, error)
	UpdateTableBranchById(*UpdateTableByIdParams) error
	DeleteTableBranchById(*DeleteTableByIdParams) error
}
