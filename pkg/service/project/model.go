package project

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetProjectListParams struct {
	Context     context.Context
	OrgId       []string
	WorkspaceId []string
}

type GetProjectByIdParams struct {
	Context   context.Context
	ProjectId string
}

type GetProjectParams struct {
	Context context.Context
	OrgId   string
	Project model.Project
}

type CreateProjectParams struct {
	Context     context.Context
	OrgId       string
	WorkspaceId string
	Project     model.Project
}

type UpdateProjectByIdParams struct {
	Context   context.Context
	ProjectId string
	Project   model.Project
}

type DeleteProjectByIdParams struct {
	Context   context.Context
	ProjectId string
}
