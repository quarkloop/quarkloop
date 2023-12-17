package project

type Service interface {
	GetProjectList(*GetProjectListParams) ([]Project, error)
	GetProjectById(*GetProjectByIdParams) (*Project, error)
	GetProject(*GetProjectParams) (*Project, error)
	CreateProject(*CreateProjectParams) (*Project, error)
	UpdateProjectById(*UpdateProjectByIdParams) error
	DeleteProjectById(*DeleteProjectByIdParams) error
}
