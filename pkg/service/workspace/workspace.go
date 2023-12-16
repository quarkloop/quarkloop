package workspace

type Service interface {
	GetWorkspaceList(*GetWorkspaceListParams) ([]Workspace, error)
	GetWorkspaceById(*GetWorkspaceByIdParams) (*Workspace, error)
	GetWorkspace(*GetWorkspaceParams) (*Workspace, error)
	CreateWorkspace(*CreateWorkspaceParams) (*Workspace, error)
	UpdateWorkspaceById(*UpdateWorkspaceByIdParams) error
	DeleteWorkspaceById(*DeleteWorkspaceByIdParams) error
}
