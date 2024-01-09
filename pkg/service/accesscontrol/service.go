package accesscontrol

import "context"

type Service interface {
	// access control query
	Evaluate(context.Context, *EvaluateQuery) error
	GetUserAccessList(context.Context, *GetUserAssignmentListQuery) ([]UserAssignment, error)
	GetUserAccessById(context.Context, *GetUserAssignmentByIdQuery) (*UserAssignment, error)

	// access control mutation
	GrantUserAccess(context.Context, *CreateUserAssignmentCommand) (*UserAssignment, error)
	UpdateUserAccessById(context.Context, *UpdateUserAssignmentByIdCommand) error
	RevokeUserAccessById(context.Context, *DeleteUserAssignmentByIdCommand) error

	// user group query
	GetUserGroupList(context.Context, *GetUserGroupListQuery) ([]UserGroup, error)
	GetUserGroupById(context.Context, *GetUserGroupByIdQuery) (*UserGroup, error)

	// user group mutation
	CreateUserGroup(context.Context, *CreateUserGroupCommand) (*UserGroup, error)
	UpdateUserGroupById(context.Context, *UpdateUserGroupByIdCommand) error
	DeleteUserGroupById(context.Context, *DeleteUserGroupByIdCommand) error

	// user role query
	GetUserRoleList(context.Context, *GetUserRoleListQuery) ([]UserRole, error)
	GetUserRoleById(context.Context, *GetUserRoleByIdQuery) (*UserRole, error)

	// user role mutation
	CreateUserRole(context.Context, *CreateUserRoleCommand) (*UserRole, error)
	//UpdateUserRoleById(context.Context, *UpdateUserRoleByIdCommand) error
	DeleteUserRoleById(context.Context, *DeleteUserRoleByIdCommand) error
}
