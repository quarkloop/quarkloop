package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type Service interface {
	// workspace
	GetWorkspaceList(*gin.Context, *GetWorkspaceListQuery) ([]*Workspace, error)
	GetWorkspaceById(*gin.Context, *GetWorkspaceByIdQuery) (*Workspace, error)
	// TODO: rewrite
	// GetWorkspace(context.Context, *GetWorkspaceQuery) (*Workspace, error)
	CreateWorkspace(*gin.Context, *CreateWorkspaceCommand) (*Workspace, error)
	UpdateWorkspaceById(*gin.Context, *UpdateWorkspaceByIdCommand) error
	DeleteWorkspaceById(*gin.Context, *DeleteWorkspaceByIdCommand) error

	// project
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)

	// user
	GetUserList(*gin.Context, *GetUserListQuery) ([]*user.User, error)
}
