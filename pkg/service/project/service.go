package project

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type Service interface {
	// query
	GetProjectById(*gin.Context, *GetProjectByIdQuery) (*Project, error)
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*Project, error)
	GetUserAssignmentList(*gin.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateProject(*gin.Context, *CreateProjectCommand) (*Project, error)
	UpdateProjectById(*gin.Context, *UpdateProjectByIdCommand) error
	DeleteProjectById(*gin.Context, *DeleteProjectByIdCommand) error
}
