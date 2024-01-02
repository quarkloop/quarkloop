package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
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
	aclService       accesscontrol.Service
	quotaService     quota.Service
}

func NewWorkspaceApi(service workspace.Service, aclService accesscontrol.Service, quotaService quota.Service) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: service,
		aclService:       aclService,
		quotaService:     quotaService,
	}
}
