package project

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type Api interface {
	GetProjectById(*gin.Context)
	CreateProject(*gin.Context)
	UpdateProjectById(*gin.Context)
	DeleteProjectById(*gin.Context)
}

type ProjectApi struct {
	projectService project.Service
}

func NewProjectApi(service project.Service) *ProjectApi {
	return &ProjectApi{
		projectService: service,
	}
}
