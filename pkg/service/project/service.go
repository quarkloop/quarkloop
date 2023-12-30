package project

import "github.com/gin-gonic/gin"

type Service interface {
	GetProjectById(*gin.Context, *GetProjectByIdQuery) (*Project, error)
	// TODO: rewrite
	//GetProject(context.Context, *GetProjectQuery) (*Project, error)
	CreateProject(*gin.Context, *CreateProjectCommand) (*Project, error)
	UpdateProjectById(*gin.Context, *UpdateProjectByIdCommand) error
	DeleteProjectById(*gin.Context, *DeleteProjectByIdCommand) error
}
