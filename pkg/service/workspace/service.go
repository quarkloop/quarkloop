package workspace

import "github.com/gin-gonic/gin"

type Service interface {
	GetWorkspaceList(*gin.Context, *GetWorkspaceListQuery) ([]*Workspace, error)
	GetWorkspaceById(*gin.Context, *GetWorkspaceByIdQuery) (*Workspace, error)
	// TODO: rewrite
	// GetWorkspace(context.Context, *GetWorkspaceQuery) (*Workspace, error)
	CreateWorkspace(*gin.Context, *CreateWorkspaceCommand) (*Workspace, error)
	UpdateWorkspaceById(*gin.Context, *UpdateWorkspaceByIdCommand) error
	DeleteWorkspaceById(*gin.Context, *DeleteWorkspaceByIdCommand) error
}
