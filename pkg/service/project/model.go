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

// GetProjectList

type GetProjectListQuery struct {
	UserId int
}

// GetProjectById

type GetProjectByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
	ProjectId   int `uri:"projectId" binding:"required"`
}

type GetProjectByIdQuery struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// type GetProjectQuery struct {
// 	OrgId   int
// 	Project Project
// }

//  CreateProject

type CreateProjectUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type CreateProjectCommand struct {
	OrgId       int
	WorkspaceId int
	Project
}

// UpdateProjectById

type UpdateProjectByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
	ProjectId   int `uri:"projectId" binding:"required"`
}

type UpdateProjectByIdCommand struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
	Project
}

// DeleteProjectById

type DeleteProjectByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
	ProjectId   int `uri:"projectId" binding:"required"`
}

type DeleteProjectByIdCommand struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
}
