package accesscontrol

import (
	"errors"
	"time"
)

var (
	ErrResourceNotFound   = errors.New("resource not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrPermissionNotFound = errors.New("permission not found")
	ErrPermissionDenied   = errors.New("permission denied")
)

const (
	DirectMembership    = 1
	InheritedMembership = 2
)

const (
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

type OrgMember struct {
	// id
	Id     int32 `json:"id"`
	UserId int32 `json:"userId"`
	RoleId int32 `json:"userRoleId"`
	OrgId  int32 `json:"orgId"`

	// membership
	ExpireDate *time.Time `json:"expireDate"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type WorkspaceMember struct {
	// id
	Id          int32 `json:"id"`
	UserId      int32 `json:"userId"`
	RoleId      int32 `json:"userRoleId"`
	WorkspaceId int32 `json:"workspaceId"`

	// membership
	Type       int32      `json:"type"`
	Source     *int32     `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type ProjectMember struct {
	// id
	Id        int32 `json:"id"`
	UserId    int32 `json:"userId"`
	RoleId    int32 `json:"userRoleId"`
	ProjectId int32 `json:"projectId"`

	// membership
	Type       int32      `json:"type"`
	Source     *int32     `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

// type UserAssignment struct {
// 	// id
// 	Id          int32  `json:"id"`
// 	OrgId       *int32 `json:"orgId"`
// 	WorkspaceId *int32 `json:"workspaceId"`
// 	ProjectId   *int32 `json:"projectId"`
// 	UserId      *int32 `json:"userId"`
// 	UserGroupId *int32 `json:"userGroupId"`
// 	RoleId  int32  `json:"userRoleId"`

// 	// history
// 	CreatedAt time.Time  `json:"createdAt"`
// 	CreatedBy string     `json:"createdBy"`
// 	UpdatedAt *time.Time `json:"updatedAt"`
// 	UpdatedBy *string    `json:"updatedBy"`
// }

// type UserGroup struct {
// 	// id
// 	Id    int32 `json:"id"`
// 	OrgId int32 `json:"orgId"`

// 	// user
// 	Name  string `json:"name"`
// 	Users []int32  `json:"users"`

// 	// history
// 	CreatedAt time.Time  `json:"createdAt"`
// 	CreatedBy string     `json:"createdBy"`
// 	UpdatedAt *time.Time `json:"updatedAt"`
// 	UpdatedBy *string    `json:"updatedBy"`
// }

type Role struct {
	// id
	Id int32 `json:"id"`

	// user
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

// Evaluate

type EvaluateQuery struct {
	UserId     int32
	Permission string

	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// GrantUserAccess

type GrantUserAccessCommand struct {
	UserId int32
	RoleId int32

	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// RevokeUserAccess

type RevokeUserAccessCommand struct {
	UserId int32
}

////////////////////////////////////////////

// GetOrgMemberList

type GetOrgMemberListQuery struct {
	OrgId int32
}

// GetOrgMemberById

type GetOrgMemberByIdQuery struct {
	OrgId    int32
	MemberId int32
}

// GetOrgMemberByUserId

type GetOrgMemberByUserIdQuery struct {
	OrgId  int32
	UserId int32
}

// GetWorkspaceMemberList

type GetWorkspaceMemberListQuery struct {
	WorkspaceId int32
}

// GetWorkspaceMemberById

type GetWorkspaceMemberByIdQuery struct {
	WorkspaceId int32
	MemberId    int32
}

// GetWorkspaceMemberByUserId

type GetWorkspaceMemberByUserIdQuery struct {
	WorkspaceId int32
	UserId      int32
}

// GetProjectMemberList

type GetProjectMemberListQuery struct {
	ProjectId int32
}

// GetProjectMemberById

type GetProjectMemberByIdQuery struct {
	ProjectId int32
	MemberId  int32
}

// GetProjectMemberByUserId

type GetProjectMemberByUserIdQuery struct {
	ProjectId int32
	UserId    int32
}

/////////////////////////////////////////////

// CreateOrgMember

type CreateOrgMemberCommand struct {
	UserId    int32
	RoleId    int32
	OrgId     int32
	CreatedBy string

	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateOrgMemberById

type UpdateOrgMemberByIdCommand struct {
	OrgId     int32
	MemberId  int32
	UpdatedBy string

	RoleId     int32      `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteOrgMemberById

type DeleteOrgMemberByIdCommand struct {
	OrgId    int32
	MemberId int32
}

// CreateWorkspaceMember

type CreateWorkspaceMemberCommand struct {
	UserId      int32
	RoleId      int32
	WorkspaceId int32
	CreatedBy   string

	Type       int32      `json:"type"`
	Source     int32      `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateWorkspaceMemberById

type UpdateWorkspaceMemberByIdCommand struct {
	WorkspaceId int32
	MemberId    int32
	UpdatedBy   string

	RoleId     int32      `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteWorkspaceMemberById

type DeleteWorkspaceMemberByIdCommand struct {
	WorkspaceId int32
	MemberId    int32
}

// CreateProjectMember

type CreateProjectMemberCommand struct {
	UserId    int32
	RoleId    int32
	ProjectId int32
	CreatedBy string

	Type       int32      `json:"type"`
	Source     int32      `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateProjectMemberById

type UpdateProjectMemberByIdCommand struct {
	ProjectId int32
	MemberId  int32
	UpdatedBy string

	RoleId     int32      `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteProjectMemberById

type DeleteProjectMemberByIdCommand struct {
	ProjectId int32
	MemberId  int32
}

/////////////////////////////////////////////

// GetRoleById

type GetRoleByIdUriParams struct {
	RoleId int32 `uri:"roleId" binding:"required"`
}

type GetRoleByIdQuery struct {
	RoleId int32
}

// CreateRole

type CreateRoleCommand struct {
	CreatedBy string

	Name string `json:"name"`
}

// UpdateRoleById

type UpdateRoleByIdUriParams struct {
	RoleId int32 `uri:"roleId" binding:"required"`
}

type UpdateRoleByIdCommand struct {
	RoleId    int32
	UpdatedBy string

	Name string `json:"name"`
}

// DeleteRoleById

type DeleteRoleByIdUriParams struct {
	RoleId int32 `uri:"roleId" binding:"required"`
}

type DeleteRoleByIdCommand struct {
	RoleId int32
}

////////////////////////////////////////////////////

// // GetUserGroupList

// type GetUserGroupListUriParams struct {
// 	OrgId int32 `uri:"orgId" binding:"required"`
// }

// type GetUserGroupListQuery struct {
// 	OrgId int32
// }

// // GetUserGroupById

// type GetUserGroupByIdUriParams struct {
// 	OrgId       int32 `uri:"orgId" binding:"required"`
// 	UserGroupId int32 `uri:"groupId" binding:"required"`
// }

// type GetUserGroupByIdQuery struct {
// 	OrgId       int32
// 	UserGroupId int32
// }

// // CreateUserGroup

// type CreateUserGroupUriParams struct {
// 	OrgId int32 `uri:"orgId" binding:"required"`
// }

// type CreateUserGroupCommand struct {
// 	OrgId     int32
// 	CreatedBy string

// 	Name  string `json:"name"`
// 	Users []int32  `json:"users"`
// }

// // UpdateUserGroupById

// type UpdateUserGroupByIdUriParams struct {
// 	OrgId       int32 `uri:"orgId" binding:"required"`
// 	UserGroupId int32 `uri:"groupId" binding:"required"`
// }

// type UpdateUserGroupByIdCommand struct {
// 	OrgId       int32
// 	UserGroupId int32
// 	UserGroup   *UserGroup
// }

// // DeleteUserGroupById

// type DeleteUserGroupByIdUriParams struct {
// 	OrgId       int32 `uri:"orgId" binding:"required"`
// 	UserGroupId int32 `uri:"groupId" binding:"required"`
// }

// type DeleteUserGroupByIdCommand struct {
// 	OrgId       int32
// 	UserGroupId int32
// }

// // GetUserAssignmentList

// type GetUserAssignmentListQuery struct {
// 	OrgId       int32
// 	WorkspaceId int32
// 	ProjectId   int32
// }

// // GetUserAssignmentById

// type GetUserAssignmentByIdQuery struct {
// 	UserAssignmentId int32
// }

// // CreateUserAssignment

// type CreateUserAssignmentCommand struct {
// 	CreatedBy string

// 	UserId      int32
// 	UserGroupId int32
// 	RoleId  int32

// 	OrgId       int32
// 	WorkspaceId int32
// 	ProjectId   int32
// }

// // UpdateUserAssignmentById

// type UpdateUserAssignmentByIdCommand struct {
// 	OrgId            int32
// 	UserAssignmentId int32
// 	Role         *UserAssignment
// }

// // DeleteUserAssignmentById

// type DeleteUserAssignmentByIdCommand struct {
// 	UserAssignmentId int32

// 	OrgId       int32
// 	WorkspaceId int32
// 	ProjectId   int32
// }

/////////

// type GetUserGroupListQuery struct {
// 	OrgId int32
// }

// type GetUserGroupByIdQuery struct {
// 	UserGroupId int32
// }

// type CreateUserGroupCommand struct {
// 	OrgId     int32
// 	UserGroup *accesscontrol.UserGroup
// }

// type UpdateUserGroupByIdCommand struct {
// 	UserGroupId int32
// 	UserGroup   *accesscontrol.UserGroup
// }

// type DeleteUserGroupByIdCommand struct {
// 	OrgId       int32
// 	UserGroupId int32
// }

// type GetRoleByIdQuery struct {
// 	RoleId int32
// }

// type CreateRoleCommand struct {
// 	OrgId    int32
// 	Role *accesscontrol.Role
// }

// type UpdateRoleByIdCommand struct {
// 	RoleId int32
// 	Role   *accesscontrol.Role
// }
