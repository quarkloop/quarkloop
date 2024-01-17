package org

import (
	"errors"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrOrgNotFound       = errors.New("org not found")
	ErrOrgMemberNotFound = errors.New("org member not found")
	ErrOrgAlreadyExists  = errors.New("org with same scopeId already exists")
)

type GetOrgListQuery struct {
	UserId     int32
	Visibility model.ScopeVisibility
}

type GetOrgByIdQuery struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}

type GetOrgVisibilityByIdQuery struct {
	OrgId int32
}

type CreateOrgCommand struct {
	CreatedBy string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility"`
}

type UpdateOrgByIdCommand struct {
	UpdatedBy string
	OrgId     int32 `uri:"orgId" binding:"required"`

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
}

type DeleteOrgByIdCommand struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}

type GetWorkspaceListQuery struct {
	OrgId      int32 `uri:"orgId" binding:"required"`
	Visibility model.ScopeVisibility
}

type GetProjectListQuery struct {
	OrgId      int32 `uri:"orgId" binding:"required"`
	Visibility model.ScopeVisibility
}

type GetMemberListQuery struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}

type GetUserAssignmentListQuery struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}
