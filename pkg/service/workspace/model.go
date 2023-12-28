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

type GetWorkspaceListParams struct {
	OrgId []int
}

type GetWorkspaceByIdParams struct {
	WorkspaceId int
}

type GetWorkspaceParams struct {
	OrgId     int
	Workspace Workspace
}

type CreateWorkspaceParams struct {
	OrgId     int
	Workspace Workspace
}

type UpdateWorkspaceByIdParams struct {
	WorkspaceId int
	Workspace   Workspace
}

type DeleteWorkspaceByIdParams struct {
	WorkspaceId int
}
