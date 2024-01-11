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
	Id     int `json:"id"`
	UserId int `json:"userId"`
	RoleId int `json:"userRoleId"`
	OrgId  int `json:"orgId"`

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
	Id          int `json:"id"`
	UserId      int `json:"userId"`
	RoleId      int `json:"userRoleId"`
	WorkspaceId int `json:"workspaceId"`

	// membership
	Type       int        `json:"type"`
	Source     *int       `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type ProjectMember struct {
	// id
	Id        int `json:"id"`
	UserId    int `json:"userId"`
	RoleId    int `json:"userRoleId"`
	ProjectId int `json:"projectId"`

	// membership
	Type       int        `json:"type"`
	Source     *int       `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

// type UserAssignment struct {
// 	// id
// 	Id          int  `json:"id"`
// 	OrgId       *int `json:"orgId"`
// 	WorkspaceId *int `json:"workspaceId"`
// 	ProjectId   *int `json:"projectId"`
// 	UserId      *int `json:"userId"`
// 	UserGroupId *int `json:"userGroupId"`
// 	RoleId  int  `json:"userRoleId"`

// 	// history
// 	CreatedAt time.Time  `json:"createdAt"`
// 	CreatedBy string     `json:"createdBy"`
// 	UpdatedAt *time.Time `json:"updatedAt"`
// 	UpdatedBy *string    `json:"updatedBy"`
// }

// type UserGroup struct {
// 	// id
// 	Id    int `json:"id"`
// 	OrgId int `json:"orgId"`

// 	// user
// 	Name  string `json:"name"`
// 	Users []int  `json:"users"`

// 	// history
// 	CreatedAt time.Time  `json:"createdAt"`
// 	CreatedBy string     `json:"createdBy"`
// 	UpdatedAt *time.Time `json:"updatedAt"`
// 	UpdatedBy *string    `json:"updatedBy"`
// }

type Role struct {
	// id
	Id int `json:"id"`

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
	UserId     int
	Permission string

	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// GrantUserAccess

type GrantUserAccessCommand struct {
	UserId int
	RoleId int

	OrgId       int
	WorkspaceId int
	ProjectId   int
}

// RevokeUserAccess

type RevokeUserAccessCommand struct {
	UserId int
}

////////////////////////////////////////////

// GetOrgMemberList

type GetOrgMemberListQuery struct {
	OrgId int
}

// GetOrgMemberById

type GetOrgMemberByIdQuery struct {
	OrgId    int
	MemberId int
}

// GetOrgMemberByUserId

type GetOrgMemberByUserIdQuery struct {
	OrgId  int
	UserId int
}

// GetWorkspaceMemberList

type GetWorkspaceMemberListQuery struct {
	WorkspaceId int
}

// GetWorkspaceMemberById

type GetWorkspaceMemberByIdQuery struct {
	WorkspaceId int
	MemberId    int
}

// GetWorkspaceMemberByUserId

type GetWorkspaceMemberByUserIdQuery struct {
	WorkspaceId int
	UserId      int
}

// GetProjectMemberList

type GetProjectMemberListQuery struct {
	ProjectId int
}

// GetProjectMemberById

type GetProjectMemberByIdQuery struct {
	ProjectId int
	MemberId  int
}

// GetProjectMemberByUserId

type GetProjectMemberByUserIdQuery struct {
	ProjectId int
	UserId    int
}

/////////////////////////////////////////////

// CreateOrgMember

type CreateOrgMemberCommand struct {
	UserId    int
	RoleId    int
	OrgId     int
	CreatedBy string

	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateOrgMemberById

type UpdateOrgMemberByIdCommand struct {
	OrgId     int
	MemberId  int
	UpdatedBy string

	RoleId     int        `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteOrgMemberById

type DeleteOrgMemberByIdCommand struct {
	OrgId    int
	MemberId int
}

// CreateWorkspaceMember

type CreateWorkspaceMemberCommand struct {
	UserId      int
	RoleId      int
	WorkspaceId int
	CreatedBy   string

	Type       int        `json:"type"`
	Source     int        `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateWorkspaceMemberById

type UpdateWorkspaceMemberByIdCommand struct {
	WorkspaceId int
	MemberId    int
	UpdatedBy   string

	RoleId     int        `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteWorkspaceMemberById

type DeleteWorkspaceMemberByIdCommand struct {
	WorkspaceId int
	MemberId    int
}

// CreateProjectMember

type CreateProjectMemberCommand struct {
	UserId    int
	RoleId    int
	ProjectId int
	CreatedBy string

	Type       int        `json:"type"`
	Source     int        `json:"source"`
	ExpireDate *time.Time `json:"expireDate"`
}

// UpdateProjectMemberById

type UpdateProjectMemberByIdCommand struct {
	ProjectId int
	MemberId  int
	UpdatedBy string

	RoleId     int        `json:"roleId"`
	ExpireDate *time.Time `json:"expireDate"`
}

// DeleteProjectMemberById

type DeleteProjectMemberByIdCommand struct {
	ProjectId int
	MemberId  int
}

/////////////////////////////////////////////

// GetRoleById

type GetRoleByIdUriParams struct {
	RoleId int `uri:"roleId" binding:"required"`
}

type GetRoleByIdQuery struct {
	RoleId int
}

// CreateRole

type CreateRoleCommand struct {
	CreatedBy string

	Name string `json:"name"`
}

// UpdateRoleById

type UpdateRoleByIdUriParams struct {
	RoleId int `uri:"roleId" binding:"required"`
}

type UpdateRoleByIdCommand struct {
	RoleId    int
	UpdatedBy string

	Name string `json:"name"`
}

// DeleteRoleById

type DeleteRoleByIdUriParams struct {
	RoleId int `uri:"roleId" binding:"required"`
}

type DeleteRoleByIdCommand struct {
	RoleId int
}

////////////////////////////////////////////////////

// // GetUserGroupList

// type GetUserGroupListUriParams struct {
// 	OrgId int `uri:"orgId" binding:"required"`
// }

// type GetUserGroupListQuery struct {
// 	OrgId int
// }

// // GetUserGroupById

// type GetUserGroupByIdUriParams struct {
// 	OrgId       int `uri:"orgId" binding:"required"`
// 	UserGroupId int `uri:"groupId" binding:"required"`
// }

// type GetUserGroupByIdQuery struct {
// 	OrgId       int
// 	UserGroupId int
// }

// // CreateUserGroup

// type CreateUserGroupUriParams struct {
// 	OrgId int `uri:"orgId" binding:"required"`
// }

// type CreateUserGroupCommand struct {
// 	OrgId     int
// 	CreatedBy string

// 	Name  string `json:"name"`
// 	Users []int  `json:"users"`
// }

// // UpdateUserGroupById

// type UpdateUserGroupByIdUriParams struct {
// 	OrgId       int `uri:"orgId" binding:"required"`
// 	UserGroupId int `uri:"groupId" binding:"required"`
// }

// type UpdateUserGroupByIdCommand struct {
// 	OrgId       int
// 	UserGroupId int
// 	UserGroup   *UserGroup
// }

// // DeleteUserGroupById

// type DeleteUserGroupByIdUriParams struct {
// 	OrgId       int `uri:"orgId" binding:"required"`
// 	UserGroupId int `uri:"groupId" binding:"required"`
// }

// type DeleteUserGroupByIdCommand struct {
// 	OrgId       int
// 	UserGroupId int
// }

// // GetUserAssignmentList

// type GetUserAssignmentListQuery struct {
// 	OrgId       int
// 	WorkspaceId int
// 	ProjectId   int
// }

// // GetUserAssignmentById

// type GetUserAssignmentByIdQuery struct {
// 	UserAssignmentId int
// }

// // CreateUserAssignment

// type CreateUserAssignmentCommand struct {
// 	CreatedBy string

// 	UserId      int
// 	UserGroupId int
// 	RoleId  int

// 	OrgId       int
// 	WorkspaceId int
// 	ProjectId   int
// }

// // UpdateUserAssignmentById

// type UpdateUserAssignmentByIdCommand struct {
// 	OrgId            int
// 	UserAssignmentId int
// 	Role         *UserAssignment
// }

// // DeleteUserAssignmentById

// type DeleteUserAssignmentByIdCommand struct {
// 	UserAssignmentId int

// 	OrgId       int
// 	WorkspaceId int
// 	ProjectId   int
// }

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

// type GetRoleByIdQuery struct {
// 	RoleId int
// }

// type CreateRoleCommand struct {
// 	OrgId    int
// 	Role *accesscontrol.Role
// }

// type UpdateRoleByIdCommand struct {
// 	RoleId int
// 	Role   *accesscontrol.Role
// }
