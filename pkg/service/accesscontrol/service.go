package accesscontrol

import "context"

type Service interface {
	// access control
	Evaluate(ctx context.Context, permission string, p *EvaluateFilterParams) (bool, error)
	ListUserAccesses(ctx context.Context, orgId int) ([]UserAssignment, error)
	GetUserAccessById(ctx context.Context, userAssignmentId int) (*UserAssignment, error)
	GrantUserAccess(ctx context.Context, orgId int, userRole *UserAssignment) (*UserAssignment, error)
	UpdateUserAccessById(ctx context.Context, userAssignmentId int, userRole *UserAssignment) error
	RevokeUserAccessById(ctx context.Context, orgId int, userAssignmentId int) error

	// user groups
	ListUserGroups(ctx context.Context, orgId int) ([]UserGroup, error)
	GetUserGroupById(ctx context.Context, userGroupId int) (*UserGroup, error)
	CreateUserGroup(ctx context.Context, orgId int, userGroup *UserGroup) (*UserGroup, error)
	UpdateUserGroupById(ctx context.Context, userGroupId int, userGroup *UserGroup) error
	DeleteUserGroupById(ctx context.Context, orgId int, userGroupId int) error

	// role
	ListUserRoles(ctx context.Context, orgId int) ([]UserRole, error)
	GetUserRoleById(ctx context.Context, userRoleId int) (*UserRole, error)
	CreateUserRole(ctx context.Context, orgId int, userRole *UserRole) (*UserRole, error)
	UpdateUserRoleById(ctx context.Context, userRoleId int, userRole *UserRole) error
	DeleteUserRoleById(ctx context.Context, orgId int, userRoleId int) error
}
