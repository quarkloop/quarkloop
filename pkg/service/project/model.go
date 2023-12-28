package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

var (
	ErrProjectNotFound = errors.New("project not found")
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
	Name        string                      `json:"name,omitempty" form:"name,omitempty"`
	Description string                      `json:"description,omitempty"`
	Visibility  *model.ScopeVisibility      `json:"visibility,omitempty" form:"visibility,omitempty"`
	Path        string                      `json:"path,omitempty"`
	Metadata    json.RawMessage             `json:"metadata,omitempty"`
	Branches    []*table_branch.TableBranch `json:"branches,omitempty"`

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
	OrgId       []int
	WorkspaceId []int
}

type GetProjectByIdParams struct {
	ProjectId int
}

type GetProjectParams struct {
	OrgId   int
	Project Project
}

type CreateProjectParams struct {
	OrgId       int
	WorkspaceId int
	Project     Project
}

type UpdateProjectByIdParams struct {
	ProjectId int
	Project   Project
}

type DeleteProjectByIdParams struct {
	ProjectId int
}
