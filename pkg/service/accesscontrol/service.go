package accesscontrol

import "context"

type Service interface {
	// access control query
	Evaluate(ctx context.Context, permission string, p *EvaluateFilterQuery) error
	ListUserAccesses(ctx context.Context, orgId int) ([]UserAssignment, error)
	GetUserAccessById(ctx context.Context, userAssignmentId int) (*UserAssignment, error)

	// access control mutation
	GrantUserAccess(ctx context.Context, orgId int, userRole *UserAssignment) (*UserAssignment, error)
	UpdateUserAccessById(ctx context.Context, userAssignmentId int, userRole *UserAssignment) error
	RevokeUserAccessById(ctx context.Context, orgId int, userAssignmentId int) error

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
	UpdateUserRoleById(context.Context, *UpdateUserRoleByIdCommand) error
	DeleteUserRoleById(context.Context, *DeleteUserRoleByIdCommand) error
}
