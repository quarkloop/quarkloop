package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Api interface {
	// workspace
	GetWorkspaceList(*gin.Context)
	GetWorkspaceById(*gin.Context)
	// TODO: rewrite
	// GetWorkspace(*gin.Context)
	CreateWorkspace(*gin.Context)
	UpdateWorkspaceById(*gin.Context)
	DeleteWorkspaceById(*gin.Context)

	// project
	GetProjectList(*gin.Context)
}

type WorkspaceApi struct {
	workspaceService workspace.Service
}

func NewWorkspaceApi(service workspace.Service) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: service,
	}
}
