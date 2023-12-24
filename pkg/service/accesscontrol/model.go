package accesscontrol

import (
	"errors"
	"time"
)

var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrUserGroupNotFound  = errors.New("user group not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrPermissionNotFound = errors.New("permission not found")
)

var (
	// org actions
	ActionOrgRead   = "org:read"
	ActionOrgCreate = "org:create"
	ActionOrgUpdate = "org:update"
	ActionOrgDelete = "org:delete"

	ActionOrgSettingsRead   = "org.settings:read"
	ActionOrgSettingsUpdate = "org.settings:update"

	ActionOrgQuotaRead   = "org.quota:read"
	ActionOrgQuotaCreate = "org.quota:create"
	ActionOrgQuotaUpdate = "org.quota:update"
	ActionOrgQuotaDelete = "org.quota:delete"

	ActionOrgUserRead   = "org.user:read"
	ActionOrgUserCreate = "org.user:create"
	ActionOrgUserUpdate = "org.user:update"
	ActionOrgUserDelete = "org.user:delete"

	// workspace actions
	ActionWorkspaceRead   = "workspace:read"
	ActionWorkspaceCreate = "workspace:create"
	ActionWorkspaceUpdate = "workspace:update"
	ActionWorkspaceDelete = "workspace:delete"

	ActionWorkspaceSettingsRead   = "workspace.settings:read"
	ActionWorkspaceSettingsUpdate = "workspace.settings:update"

	ActionWorkspaceQuotaRead   = "workspace.quota:read"
	ActionWorkspaceQuotaCreate = "workspace.quota:create"
	ActionWorkspaceQuotaUpdate = "workspace.quota:update"
	ActionWorkspaceQuotaDelete = "workspace.quota:delete"

	ActionWorkspaceUserRead   = "workspace.user:read"
	ActionWorkspaceUserCreate = "workspace.user:create"
	ActionWorkspaceUserUpdate = "workspace.user:update"
	ActionWorkspaceUserDelete = "workspace.user:delete"

	// project actions
	ActionProjectRead   = "project:read"
	ActionProjectCreate = "project:create"
	ActionProjectUpdate = "project:update"
	ActionProjectDelete = "project:delete"

	ActionProjectSettingsRead   = "project.settings:read"
	ActionProjectSettingsUpdate = "project.settings:update"

	ActionProjectQuotaRead   = "project.quota:read"
	ActionProjectQuotaCreate = "project.quota:create"
	ActionProjectQuotaUpdate = "project.quota:update"
	ActionProjectQuotaDelete = "project.quota:delete"

	ActionProjectUserRead   = "project.user:read"
	ActionProjectUserCreate = "project.user:create"
	ActionProjectUserUpdate = "project.user:update"
	ActionProjectUserDelete = "project.user:delete"
)

type UserAssignment struct {
	// id
	Id          int `json:"id"`
	OrgId       int `json:"orgId"`
	WorkspaceId int `json:"workspaceId"`
	ProjectId   int `json:"projectId"`
	UserGroupId int `json:"userGroupId"`
	UserRoleId  int `json:"userRoleId"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

type UserGroup struct {
	// id
	Id     int `json:"id"`
	OrgId  int `json:"orgId"`
	UserId int `json:"userId"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

type UserRole struct {
	// id
	Id    int `json:"id"`
	OrgId int `json:"orgId"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

type Permission struct {
	// id
	Id int `json:"id"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}
