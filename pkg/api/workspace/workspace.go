package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Api interface {
	// query
	GetWorkspaceById(*gin.Context)
	GetWorkspaceList(*gin.Context)
	GetProjectList(*gin.Context)
	GetUserList(*gin.Context)

	// mutation
	CreateWorkspace(*gin.Context)
	UpdateWorkspaceById(*gin.Context)
	DeleteWorkspaceById(*gin.Context)
}

type WorkspaceApi struct {
	workspaceService workspace.Service
}

func NewWorkspaceApi(service workspace.Service) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: service,
	}
}
