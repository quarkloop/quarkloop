package accesscontrol

import "context"

type Service interface {
	// user access query and mutation
	EvaluateUserAccess(context.Context, *EvaluateQuery) (bool, error)
	GrantUserAccess(context.Context, *GrantUserAccessCommand) error
	RevokeUserAccess(context.Context, *RevokeUserAccessCommand) error

	// member query
	GetOrgMemberList(context.Context, *GetOrgMemberListQuery) ([]*Member, error)
	GetWorkspaceMemberList(context.Context, *GetWorkspaceMemberListQuery) ([]*Member, error)
	GetProjectMemberList(context.Context, *GetProjectMemberListQuery) ([]*Member, error)

	// resource query
	GetOrgList(context.Context, *GetOrgListQuery) ([]int32, error)
	GetWorkspaceList(context.Context, *GetWorkspaceListQuery) ([]int32, error)
	GetProjectList(context.Context, *GetProjectListQuery) ([]int32, error)

	// member mutation
	// CreateOrgMember(context.Context, *CreateOrgMemberCommand) (*OrgMember, error)
	// UpdateOrgMemberById(context.Context, *UpdateOrgMemberByIdCommand) error
	// DeleteOrgMemberById(context.Context, *DeleteOrgMemberByIdCommand) error
	// CreateWorkspaceMember(context.Context, *CreateWorkspaceMemberCommand) (*WorkspaceMember, error)
	// UpdateWorkspaceMemberById(context.Context, *UpdateWorkspaceMemberByIdCommand) error
	// DeleteWorkspaceMemberById(context.Context, *DeleteWorkspaceMemberByIdCommand) error
	// CreateProjectMember(context.Context, *CreateProjectMemberCommand) (*ProjectMember, error)
	// UpdateProjectMemberById(context.Context, *UpdateProjectMemberByIdCommand) error
	// DeleteProjectMemberById(context.Context, *DeleteProjectMemberByIdCommand) error

	// user role query and mutation
	// GetRoleList(context.Context) ([]*Role, error)
	// GetRoleById(context.Context, *GetRoleByIdQuery) (*Role, error)
	// CreateRole(context.Context, *CreateRoleCommand) (*Role, error)
	// DeleteRoleById(context.Context, *DeleteRoleByIdCommand) error
}

// type UserGroupService interface {
// 	// user group query
// 	GetUserGroupList(context.Context, *GetUserGroupListQuery) ([]UserGroup, error)
// 	GetUserGroupById(context.Context, *GetUserGroupByIdQuery) (*UserGroup, error)
// 	CreateUserGroup(context.Context, *CreateUserGroupCommand) (*UserGroup, error)
// 	UpdateUserGroupById(context.Context, *UpdateUserGroupByIdCommand) error
// 	DeleteUserGroupById(context.Context, *DeleteUserGroupByIdCommand) error
// }
