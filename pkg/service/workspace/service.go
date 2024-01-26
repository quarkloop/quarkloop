package workspace

import (
	// "github.com/gin-gonic/gin"
	// "github.com/quarkloop/quarkloop/pkg/model"
	// "github.com/quarkloop/quarkloop/pkg/service/project"
	// "github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/v1/system/workspace"
)

type Service workspace.WorkspaceServiceServer

// type Service interface {
// 	// query
// 	GetWorkspaceById(*gin.Context, *GetWorkspaceByIdQuery) (*Workspace, error)
// 	GetWorkspaceVisibilityById(*gin.Context, *GetWorkspaceVisibilityByIdQuery) (model.ScopeVisibility, error)
// 	GetWorkspaceList(*gin.Context, *GetWorkspaceListQuery) ([]*Workspace, error)
// 	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)
// 	GetUserAssignmentList(*gin.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

// 	// mutation
// 	CreateWorkspace(*gin.Context, *CreateWorkspaceCommand) (*Workspace, error)
// 	UpdateWorkspaceById(*gin.Context, *UpdateWorkspaceByIdCommand) error
// 	DeleteWorkspaceById(*gin.Context, *DeleteWorkspaceByIdCommand) error
// }
