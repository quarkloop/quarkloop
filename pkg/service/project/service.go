package project

import (
	// "github.com/gin-gonic/gin"
	// "github.com/quarkloop/quarkloop/pkg/model"
	// "github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/v1/system/project"
)

type Service project.ProjectServiceServer

// type Service interface {
// 	// query
// 	GetProjectById(*gin.Context, *GetProjectByIdQuery) (*model.Project, error)
// 	GetProjectVisibilityById(*gin.Context, *GetProjectVisibilityByIdQuery) (model.ScopeVisibility, error)
// 	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*model.Project, error)
// 	GetUserAssignmentList(*gin.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

// 	// mutation
// 	CreateProject(*gin.Context, *CreateProjectCommand) (*model.Project, error)
// 	UpdateProjectById(*gin.Context, *UpdateProjectByIdCommand) error
// 	DeleteProjectById(*gin.Context, *DeleteProjectByIdCommand) error
// }
