package project

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type Api interface {
	GetProjectList(c *gin.Context)
	GetProjectById(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProjectById(c *gin.Context)
	DeleteProjectById(c *gin.Context)
}

type ProjectApi struct {
	projectService project.Service
}

func NewProjectApi(service project.Service) *ProjectApi {
	return &ProjectApi{
		projectService: service,
	}
}
