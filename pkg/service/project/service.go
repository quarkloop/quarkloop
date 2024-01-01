package project

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type Service interface {
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*Project, error)
	GetProjectById(*gin.Context, *GetProjectByIdQuery) (*Project, error)
	// TODO: rewrite
	//GetProject(context.Context, *GetProjectQuery) (*Project, error)
	CreateProject(*gin.Context, *CreateProjectCommand) (*Project, error)
	UpdateProjectById(*gin.Context, *UpdateProjectByIdCommand) error
	DeleteProjectById(*gin.Context, *DeleteProjectByIdCommand) error

	// user
	GetUserList(*gin.Context, *GetUserListQuery) ([]*user.User, error)
}
