package workspace

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetWorkspaceListParams struct {
	Context context.Context
	OrgId   []int
}

type GetWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId int
}

type GetWorkspaceParams struct {
	Context   context.Context
	OrgId     int
	Workspace model.Workspace
}

type CreateWorkspaceParams struct {
	Context   context.Context
	OrgId     int
	Workspace model.Workspace
}

type UpdateWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId int
	Workspace   model.Workspace
}

type DeleteWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId int
}
