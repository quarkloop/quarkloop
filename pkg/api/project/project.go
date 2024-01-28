package project

import (
	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/project"
)

type Api interface {
	// query
	GetProjectById(*gin.Context)
	GetProjectList(*gin.Context)
	GetMemberList(*gin.Context)

	// mutation
	CreateProject(*gin.Context)
	UpdateProjectById(*gin.Context)
	DeleteProjectById(*gin.Context)
}

type ProjectApi struct {
	projectService grpc.ProjectServiceClient

	userService   user.Service
	aclService    accesscontrol.Service
	quotaService  quota.Service
	branchService table_branch.Service
}

func NewProjectApi(
	projectService grpc.ProjectServiceClient,
	userService user.Service,
	aclService accesscontrol.Service,
	quotaService quota.Service,
	branchService table_branch.Service,
) *ProjectApi {
	return &ProjectApi{
		projectService: projectService,
		userService:    userService,
		aclService:     aclService,
		quotaService:   quotaService,
		branchService:  branchService,
	}
}

func transformGrpcSlice(slice []*system.Project) []*system.Project {
	if slice == nil {
		return []*system.Project{}
	}
	return slice
}
