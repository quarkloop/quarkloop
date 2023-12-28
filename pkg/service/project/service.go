package project

import "github.com/gin-gonic/gin"

type Service interface {
	GetProjectList(*gin.Context, *GetProjectListParams) ([]Project, error)
	GetProjectById(*gin.Context, *GetProjectByIdParams) (*Project, error)
	// TODO: rewrite
	//GetProject(context.Context, *GetProjectParams) (*Project, error)
	CreateProject(*gin.Context, *CreateProjectParams) (*Project, error)
	UpdateProjectById(*gin.Context, *UpdateProjectByIdParams) error
	DeleteProjectById(*gin.Context, *DeleteProjectByIdParams) error
}
