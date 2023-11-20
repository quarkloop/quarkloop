package workspace

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetWorkspaceList(*GetWorkspaceListParams) ([]model.Workspace, error)
	GetWorkspaceById(*GetWorkspaceByIdParams) (*model.Workspace, error)
	GetWorkspace(*GetWorkspaceParams) (*model.Workspace, error)
	CreateWorkspace(*CreateWorkspaceParams) (*model.Workspace, error)
	UpdateWorkspaceById(*UpdateWorkspaceByIdParams) error
	DeleteWorkspaceById(*DeleteWorkspaceByIdParams) error
}
