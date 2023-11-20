package project

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetProjectList(*GetProjectListParams) ([]model.Project, error)
	GetProjectById(*GetProjectByIdParams) (*model.Project, error)
	GetProject(*GetProjectParams) (*model.Project, error)
	CreateProject(*CreateProjectParams) (*model.Project, error)
	UpdateProjectById(*UpdateProjectByIdParams) error
	DeleteProjectById(*DeleteProjectByIdParams) error
}
