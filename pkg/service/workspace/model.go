package workspace

import (
	"context"
	"fmt"
	"time"
)

type Workspace struct {
	// id
	Id          int    `json:"id" form:"id"`
	ScopedId    string `json:"sid"`
	OrgId       int    `json:"orgId"`
	OrgScopedId string `json:"orgScopedId"`

	// data
	Name        string `json:"name,omitempty" form:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AccessType  *int   `json:"accessType,omitempty" form:"accessType,omitempty"`
	Path        string `json:"path,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

func (w *Workspace) GeneratePath() {
	w.Path = fmt.Sprintf("/org/%s/%s", w.OrgScopedId, w.ScopedId)
}

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
	Workspace Workspace
}

type CreateWorkspaceParams struct {
	Context   context.Context
	OrgId     int
	Workspace Workspace
}

type UpdateWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId int
	Workspace   Workspace
}

type DeleteWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId int
}
