package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Api interface {
	GetWorkspaceList(c *gin.Context)
	GetWorkspaceById(c *gin.Context)
	// TODO: rewrite
	// GetWorkspace(c *gin.Context)
	CreateWorkspace(c *gin.Context)
	UpdateWorkspaceById(c *gin.Context)
	DeleteWorkspaceById(c *gin.Context)
}

type WorkspaceApi struct {
	workspaceService workspace.Service
}

func NewWorkspaceApi(service workspace.Service) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: service,
	}
}
