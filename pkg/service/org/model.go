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
	Id      int    `json:"id"`
	ScopeId string `json:"sid"`

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

func (o *Org) GeneratePath() {
	o.Path = fmt.Sprintf("/org/%s", o.ScopeId)
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

// GetOrgVisibilityById
type GetOrgVisibilityByIdQuery struct {
	OrgId int
}

// CreateOrg

type CreateOrgCommand struct {
	CreatedBy string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility"`
}

// UpdateOrgById

type UpdateOrgByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type UpdateOrgByIdCommand struct {
	UpdatedBy string
	OrgId     int

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
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

// GetMemberList

type GetMemberListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetMemberListQuery struct {
	OrgId int
}

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId int
}
