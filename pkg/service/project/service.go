package project

import "context"

type Service interface {
	GetProjectList(context.Context, *GetProjectListParams) ([]Project, error)
	GetProjectById(context.Context, *GetProjectByIdParams) (*Project, error)
	GetProject(context.Context, *GetProjectParams) (*Project, error)
	CreateProject(context.Context, *CreateProjectParams) (*Project, error)
	UpdateProjectById(context.Context, *UpdateProjectByIdParams) error
	DeleteProjectById(context.Context, *DeleteProjectByIdParams) error
}
