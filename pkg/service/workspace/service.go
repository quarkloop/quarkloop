package workspace

import "github.com/gin-gonic/gin"

type Service interface {
	GetWorkspaceList(*gin.Context, *GetWorkspaceListParams) ([]Workspace, error)
	GetWorkspaceById(*gin.Context, *GetWorkspaceByIdParams) (*Workspace, error)
	// TODO: rewrite
	// GetWorkspace(context.Context, *GetWorkspaceParams) (*Workspace, error)
	CreateWorkspace(*gin.Context, *CreateWorkspaceParams) (*Workspace, error)
	UpdateWorkspaceById(*gin.Context, *UpdateWorkspaceByIdParams) error
	DeleteWorkspaceById(*gin.Context, *DeleteWorkspaceByIdParams) error
}
