package project_service

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetProjectServiceListParams struct {
	Context   context.Context
	ProjectId string
}

type GetProjectServiceByIdParams struct {
	Context          context.Context
	ProjectId        string
	ProjectServiceId string
}

type CreateProjectServiceParams struct {
	Context        context.Context
	ProjectId      string
	ProjectService *model.ProjectService
}

type CreateBulkProjectServiceParams struct {
	Context            context.Context
	ProjectId          string
	ProjectServiceList []model.ProjectService
}

type UpdateProjectServiceByIdParams struct {
	Context          context.Context
	ProjectServiceId string
	ProjectService   *model.ProjectService
}

type DeleteProjectServiceByIdParams struct {
	Context          context.Context
	ProjectServiceId string
}
