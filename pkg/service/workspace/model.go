package workspace

import (
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrWorkspaceNotFound = errors.New("workspace not found")
)

type Workspace struct {
	// id
	Id         int    `json:"id"`
	ScopeId    string `json:"sid"`
	OrgId      int    `json:"orgId"`
	OrgScopeId string `json:"orgScopeId"`

	// data
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Visibility  model.ScopeVisibility `json:"visibility"`
	Path        string                `json:"path"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

func (w *Workspace) GeneratePath() {
	w.Path = fmt.Sprintf("/org/%s/%s", w.OrgScopeId, w.ScopeId)
}

// GetWorkspaceList

type GetWorkspaceListQuery struct {
	UserId     int
	Visibility model.ScopeVisibility
}

// GetWorkspaceById

type GetWorkspaceByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type GetWorkspaceByIdQuery struct {
	OrgId       int
	WorkspaceId int
}

// CreateWorkspace

type CreateWorkspaceUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type CreateWorkspaceCommand struct {
	OrgId     int
	CreatedBy string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility"`
}

// UpdateWorkspaceById

type UpdateWorkspaceByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type UpdateWorkspaceByIdCommand struct {
	OrgId       int
	WorkspaceId int
	UpdatedBy   string

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
}

// DeleteWorkspaceById

type DeleteWorkspaceByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type DeleteWorkspaceByIdCommand struct {
	OrgId       int
	WorkspaceId int
}

// GetProjectList

type GetProjectListUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type GetProjectListQuery struct {
	OrgId       int
	WorkspaceId int
	Visibility  model.ScopeVisibility
}

// GetMemberList

type GetMemberListUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type GetMemberListQuery struct {
	OrgId       int
	WorkspaceId int
}

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId       int
	WorkspaceId int
}
