package project

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type Api interface {
	// query
	GetProjectById(*gin.Context)
	GetProjectList(*gin.Context)
	GetUserList(*gin.Context)

	// mutation
	CreateProject(*gin.Context)
	UpdateProjectById(*gin.Context)
	DeleteProjectById(*gin.Context)
}

type ProjectApi struct {
	projectService project.Service

	aclService    accesscontrol.Service
	quotaService  quota.Service
	branchService table_branch.Service
}

func NewProjectApi(
	service project.Service,
	aclService accesscontrol.Service,
	quotaService quota.Service,
	branchService table_branch.Service,
) *ProjectApi {
	return &ProjectApi{
		projectService: service,
		aclService:     aclService,
		quotaService:   quotaService,
		branchService:  branchService,
	}
}
