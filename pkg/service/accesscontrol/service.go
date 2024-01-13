package accesscontrol

import "context"

type Service interface {
	// user access query and mutation
	EvaluateUserAccess(context.Context, *EvaluateQuery) error
	GrantUserAccess(context.Context, *GrantUserAccessCommand) (bool, error)
	RevokeUserAccess(context.Context, *RevokeUserAccessCommand) error

	// member query
	GetOrgMemberList(context.Context, *GetOrgMemberListQuery) ([]*OrgMember, error)
	GetOrgMemberById(context.Context, *GetOrgMemberByIdQuery) (*OrgMember, error)
	GetOrgMemberByUserId(context.Context, *GetOrgMemberByUserIdQuery) (*OrgMember, error)
	GetWorkspaceMemberList(context.Context, *GetWorkspaceMemberListQuery) ([]*WorkspaceMember, error)
	GetWorkspaceMemberById(context.Context, *GetWorkspaceMemberByIdQuery) (*WorkspaceMember, error)
	GetWorkspaceMemberByUserId(context.Context, *GetWorkspaceMemberByUserIdQuery) (*WorkspaceMember, error)
	GetProjectMemberList(context.Context, *GetProjectMemberListQuery) ([]*ProjectMember, error)
	GetProjectMemberById(context.Context, *GetProjectMemberByIdQuery) (*ProjectMember, error)
	GetProjectMemberByUserId(context.Context, *GetProjectMemberByUserIdQuery) (*ProjectMember, error)

	// member mutation
	CreateOrgMember(context.Context, *CreateOrgMemberCommand) (*OrgMember, error)
	UpdateOrgMemberById(context.Context, *UpdateOrgMemberByIdCommand) error
	DeleteOrgMemberById(context.Context, *DeleteOrgMemberByIdCommand) error
	CreateWorkspaceMember(context.Context, *CreateWorkspaceMemberCommand) (*WorkspaceMember, error)
	UpdateWorkspaceMemberById(context.Context, *UpdateWorkspaceMemberByIdCommand) error
	DeleteWorkspaceMemberById(context.Context, *DeleteWorkspaceMemberByIdCommand) error
	CreateProjectMember(context.Context, *CreateProjectMemberCommand) (*ProjectMember, error)
	UpdateProjectMemberById(context.Context, *UpdateProjectMemberByIdCommand) error
	DeleteProjectMemberById(context.Context, *DeleteProjectMemberByIdCommand) error

	// user role query and mutation
	GetRoleList(context.Context) ([]*Role, error)
	GetRoleById(context.Context, *GetRoleByIdQuery) (*Role, error)
	CreateRole(context.Context, *CreateRoleCommand) (*Role, error)
	DeleteRoleById(context.Context, *DeleteRoleByIdCommand) error
}

// type UserGroupService interface {
// 	// user group query
// 	GetUserGroupList(context.Context, *GetUserGroupListQuery) ([]UserGroup, error)
// 	GetUserGroupById(context.Context, *GetUserGroupByIdQuery) (*UserGroup, error)
// 	CreateUserGroup(context.Context, *CreateUserGroupCommand) (*UserGroup, error)
// 	UpdateUserGroupById(context.Context, *UpdateUserGroupByIdCommand) error
// 	DeleteUserGroupById(context.Context, *DeleteUserGroupByIdCommand) error
// }
