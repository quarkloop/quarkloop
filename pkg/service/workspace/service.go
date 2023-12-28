package workspace

import "context"

type Service interface {
	GetWorkspaceList(context.Context, *GetWorkspaceListParams) ([]Workspace, error)
	GetWorkspaceById(context.Context, *GetWorkspaceByIdParams) (*Workspace, error)
	// TODO: rewrite
	// GetWorkspace(context.Context, *GetWorkspaceParams) (*Workspace, error)
	CreateWorkspace(context.Context, *CreateWorkspaceParams) (*Workspace, error)
	UpdateWorkspaceById(context.Context, *UpdateWorkspaceByIdParams) error
	DeleteWorkspaceById(context.Context, *DeleteWorkspaceByIdParams) error
}
