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
	ErrPermissionDenied   = errors.New("permission denied")
)

var (
	GlobalOrgId = 0

	// org actions
	ActionOrgList   = "org:list"
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
	ActionWorkspaceList   = "workspace:list"
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
	ActionProjectList   = "project:list"
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
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type UserGroup struct {
	// id
	Id     int `json:"id"`
	OrgId  int `json:"orgId"`
	UserId int `json:"userId"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type UserRole struct {
	// id
	Id    int `json:"id"`
	OrgId int `json:"orgId"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type Permission struct {
	// id
	Id int `json:"id"`

	// user
	Name string `json:"name,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type EvaluateQuery struct {
	Permission  string
	UserId      int
	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// GetUserGroupList

type GetUserGroupListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetUserGroupListQuery struct {
	OrgId int
}

// GetUserGroupById

type GetUserGroupByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	UserGroupId int `uri:"groupId" binding:"required"`
}

type GetUserGroupByIdQuery struct {
	OrgId       int
	UserGroupId int
}

// CreateUserGroup

type CreateUserGroupUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type CreateUserGroupCommand struct {
	OrgId     int
	UserGroup *UserGroup
}

// UpdateUserGroupById

type UpdateUserGroupByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	UserGroupId int `uri:"groupId" binding:"required"`
}

type UpdateUserGroupByIdCommand struct {
	OrgId       int
	UserGroupId int
	UserGroup   *UserGroup
}

// DeleteUserGroupById

type DeleteUserGroupByIdUriParams struct {
	OrgId       int `uri:"orgId" binding:"required"`
	UserGroupId int `uri:"groupId" binding:"required"`
}

type DeleteUserGroupByIdCommand struct {
	OrgId       int
	UserGroupId int
}

// GetUserRoleList

type GetUserRoleListUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type GetUserRoleListQuery struct {
	OrgId int
}

// GetUserRoleById

type GetUserRoleByIdUriParams struct {
	OrgId      int `uri:"orgId" binding:"required"`
	UserRoleId int `uri:"roleId" binding:"required"`
}

type GetUserRoleByIdQuery struct {
	OrgId      int
	UserRoleId int
}

// CreateUserRole

type CreateUserRoleUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type CreateUserRoleCommand struct {
	OrgId    int
	UserRole *UserRole
}

// UpdateUserRoleById

type UpdateUserRoleByIdUriParams struct {
	OrgId      int `uri:"orgId" binding:"required"`
	UserRoleId int `uri:"roleId" binding:"required"`
}

type UpdateUserRoleByIdCommand struct {
	OrgId      int
	UserRoleId int
	UserRole   *UserRole
}

// DeleteUserRoleById

type DeleteUserRoleByIdUriParams struct {
	OrgId      int `uri:"orgId" binding:"required"`
	UserRoleId int `uri:"roleId" binding:"required"`
}

type DeleteUserRoleByIdCommand struct {
	OrgId      int
	UserRoleId int
}

////////////////////////////////////////////////////

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId int
}

// GetUserAssignmentById

type GetUserAssignmentByIdQuery struct {
	UserAssignmentId int
}

// CreateUserAssignment

type CreateUserAssignmentCommand struct {
	OrgId    int
	UserRole *UserAssignment
}

// UpdateUserAssignmentById

type UpdateUserAssignmentByIdCommand struct {
	OrgId            int
	UserAssignmentId int
	UserRole         *UserAssignment
}

// DeleteUserAssignmentById

type DeleteUserAssignmentByIdCommand struct {
	OrgId            int
	UserAssignmentId int
}

/////////

// type GetUserGroupListQuery struct {
// 	OrgId int
// }

// type GetUserGroupByIdQuery struct {
// 	UserGroupId int
// }

// type CreateUserGroupCommand struct {
// 	OrgId     int
// 	UserGroup *accesscontrol.UserGroup
// }

// type UpdateUserGroupByIdCommand struct {
// 	UserGroupId int
// 	UserGroup   *accesscontrol.UserGroup
// }

// type DeleteUserGroupByIdCommand struct {
// 	OrgId       int
// 	UserGroupId int
// }

// type GetUserRoleByIdQuery struct {
// 	UserRoleId int
// }

// type CreateUserRoleCommand struct {
// 	OrgId    int
// 	UserRole *accesscontrol.UserRole
// }

// type UpdateUserRoleByIdCommand struct {
// 	UserRoleId int
// 	UserRole   *accesscontrol.UserRole
// }
