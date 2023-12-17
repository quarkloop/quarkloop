package table_branch

type Service interface {
	ListTableBranches(*GetTableBranchListParams) ([]TableBranch, error)
	GetTableBranchById(*GetTableBranchByIdParams) (*TableBranch, error)
	CreateTableBranch(*CreateTableBranchParams) (*TableBranch, error)
	UpdateTableBranchById(*UpdateTableBranchByIdParams) error
	DeleteTableBranchById(*DeleteTableBranchByIdParams) error
}
