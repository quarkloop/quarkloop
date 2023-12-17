package project

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Project struct {
	// id
	Id                int    `json:"id" form:"id"`
	ScopedId          string `json:"sid"`
	WorkspaceId       int    `json:"workspaceId"`
	WorkspaceScopedId string `json:"workspaceScopedId"`
	OrgId             int    `json:"orgId"`
	OrgScopedId       string `json:"orgScopedId"`

	// data
	Name        string          `json:"name,omitempty" form:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	AccessType  *int            `json:"accessType,omitempty" form:"accessType,omitempty"`
	Path        string          `json:"path,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

func (p *Project) GeneratePath() {
	p.Path = fmt.Sprintf("/org/%s/%s/%s", p.OrgScopedId, p.WorkspaceScopedId, p.ScopedId)
}

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
	Project Project
}

type CreateProjectParams struct {
	Context     context.Context
	OrgId       int
	WorkspaceId int
	Project     Project
}

type UpdateProjectByIdParams struct {
	Context   context.Context
	ProjectId int
	Project   Project
}

type DeleteProjectByIdParams struct {
	Context   context.Context
	ProjectId int
}
