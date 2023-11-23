package project_service

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project_service"
)

type Api interface {
	GetProjectServiceList(c *gin.Context)
	GetProjectServiceById(c *gin.Context)
	CreateProjectService(c *gin.Context)
	UpdateProjectServiceById(c *gin.Context)
	DeleteProjectServiceById(c *gin.Context)
}

type ProjectServiceApi struct {
	projectService project_service.Service
}

func NewProjectServiceApi(service project_service.Service) *ProjectServiceApi {
	return &ProjectServiceApi{
		projectService: service,
	}
}
