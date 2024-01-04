package project

import (
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
	Id               int    `json:"id"`
	ScopeId          string `json:"sid"`
	WorkspaceId      int    `json:"workspaceId"`
	WorkspaceScopeId string `json:"workspaceScopeId"`
	OrgId            int    `json:"orgId"`
	OrgScopeId       string `json:"orgScopeId"`

	// project
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Visibility  model.ScopeVisibility       `json:"visibility"`
	Path        string                      `json:"path"`
	Branches    []*table_branch.TableBranch `json:"branches"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

func (p *Project) GeneratePath() {
	p.Path = fmt.Sprintf("/org/%s/%s/%s", p.OrgScopeId, p.WorkspaceScopeId, p.ScopeId)
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

// GetProjectVisibilityById

type GetProjectVisibilityByIdQuery struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// GetProjectList

type GetProjectListQuery struct {
	UserId     int
	Visibility model.ScopeVisibility
}

//  CreateProject

type CreateProjectUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type CreateProjectCommand struct {
	OrgId       int
	WorkspaceId int
	CreatedBy   string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Visibility  model.ScopeVisibility `json:"visibility"`
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
	UpdatedBy   string

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
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

// GetMemberList

type GetMemberListUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
	ProjectId   int `uri:"projectId" binding:"required"`
}

type GetMemberListQuery struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId       int
	WorkspaceId int
	ProjectId   int
}
