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
	Id          int    `json:"id" form:"id"`
	ScopedId    string `json:"sid"`
	OrgId       int    `json:"orgId"`
	OrgScopedId string `json:"orgScopedId"`

	// data
	Name        string                 `json:"name,omitempty" form:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Visibility  *model.ScopeVisibility `json:"visibility,omitempty" form:"visibility,omitempty"`
	Path        string                 `json:"path,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

func (w *Workspace) GeneratePath() {
	w.Path = fmt.Sprintf("/org/%s/%s", w.OrgScopedId, w.ScopedId)
}

// GetWorkspaceList

type GetWorkspaceListQuery struct {
	UserId int
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

// GetWorkspace

// type GetWorkspaceQuery struct {
// 	OrgId     int
// 	Workspace Workspace
// }

// CreateWorkspace

type CreateWorkspaceUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type CreateWorkspaceCommand struct {
	OrgId int
	Workspace
}

// UpdateWorkspaceById

type UpdateWorkspaceByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type UpdateWorkspaceByIdCommand struct {
	OrgId       int
	WorkspaceId int
	Workspace
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
}
