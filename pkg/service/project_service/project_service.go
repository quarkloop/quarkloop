package project_service

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetProjectServiceList(*GetProjectServiceListParams) ([]model.ProjectService, error)
	GetProjectServiceById(*GetProjectServiceByIdParams) (*model.ProjectService, error)
	CreateProjectService(*CreateProjectServiceParams) (*model.ProjectService, error)
	CreateBulkProjectService(*CreateBulkProjectServiceParams) (int64, error)
	UpdateProjectServiceById(*UpdateProjectServiceByIdParams) error
	DeleteProjectServiceById(*DeleteProjectServiceByIdParams) error
}
