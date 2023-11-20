package workspace

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetWorkspaceListParams struct {
	Context context.Context
	OrgId   []string
}

type GetWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId string
}

type GetWorkspaceParams struct {
	Context   context.Context
	OrgId     string
	Workspace model.Workspace
}

type CreateWorkspaceParams struct {
	Context   context.Context
	OrgId     string
	Workspace model.Workspace
}

type UpdateWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId string
	Workspace   model.Workspace
}

type DeleteWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId string
}
