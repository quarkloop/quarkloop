package workspace

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Api interface {
	// query
	GetWorkspaceById(*gin.Context)
	GetWorkspaceList(*gin.Context)
	GetProjectList(*gin.Context)
	GetMemberList(*gin.Context)

	// mutation
	CreateWorkspace(*gin.Context)
	UpdateWorkspaceById(*gin.Context)
	DeleteWorkspaceById(*gin.Context)
}

type WorkspaceApi struct {
	workspaceService workspace.Service

	userService  user.Service
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewWorkspaceApi(
	service workspace.Service,
	userService user.Service,
	aclService accesscontrol.Service,
	quotaService quota.Service,
) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: service,
		userService:      userService,
		aclService:       aclService,
		quotaService:     quotaService,
	}
}
