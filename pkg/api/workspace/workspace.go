package workspace

import (
	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/workspace"
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
	workspaceService grpc.WorkspaceServiceClient

	userService  user.Service
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewWorkspaceApi(
	workspaceService grpc.WorkspaceServiceClient,
	userService user.Service,
	aclService accesscontrol.Service,
	quotaService quota.Service,
) *WorkspaceApi {
	return &WorkspaceApi{
		workspaceService: workspaceService,
		userService:      userService,
		aclService:       aclService,
		quotaService:     quotaService,
	}
}

func transformGrpcSlice(slice []*system.Workspace) []*system.Workspace {
	if slice == nil {
		return []*system.Workspace{}
	}
	return slice
}
