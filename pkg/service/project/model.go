package project

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetProjectListParams struct {
	Context     context.Context
	OrgId       []int
	WorkspaceId []int
}

type GetProjectByIdParams struct {
	Context   context.Context
	ProjectId int
}

type GetProjectParams struct {
	Context context.Context
	OrgId   int
	Project model.Project
}

type CreateProjectParams struct {
	Context     context.Context
	OrgId       int
	WorkspaceId int
	Project     model.Project
}

type UpdateProjectByIdParams struct {
	Context   context.Context
	ProjectId int
	Project   model.Project
}

type DeleteProjectByIdParams struct {
	Context   context.Context
	ProjectId int
}
