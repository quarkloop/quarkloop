package org

import (
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrOrgNotFound = errors.New("org not found")
)

type Org struct {
	// id
	Id       int    `json:"id" form:"id"`
	ScopedId string `json:"sid"`

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

func (o *Org) GeneratePath() {
	o.Path = fmt.Sprintf("/org/%s", o.ScopedId)
}

// GetOrgList

type GetOrgListQuery struct {
	UserId     int
	Visibility model.ScopeVisibility
}

// GetOrgById

type GetOrgByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetOrgByIdQuery struct {
	OrgId int
}

// GetOrg

// type GetOrgParams struct {
// 	Org Org
// }

// CreateOrg

type CreateOrgCommand struct {
	Org
}

// UpdateOrgById

type UpdateOrgByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type UpdateOrgByIdCommand struct {
	OrgId int
	Org
}

// DeleteOrgById

type DeleteOrgByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type DeleteOrgByIdCommand struct {
	OrgId int
}

// GetWorkspaceList

type GetWorkspaceListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetWorkspaceListQuery struct {
	OrgId      int
	Visibility model.ScopeVisibility
}

// GetProjectList

type GetProjectListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetProjectListQuery struct {
	OrgId      int
	Visibility model.ScopeVisibility
}

// GetUserList

type GetUserListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetUserListQuery struct {
	OrgId int
}
