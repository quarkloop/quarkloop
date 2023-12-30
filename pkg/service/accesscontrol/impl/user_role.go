package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) GetUserRoleList(ctx context.Context, query *accesscontrol.GetUserRoleListQuery) ([]accesscontrol.UserRole, error) {
	urList, err := s.store.ListUserRoles(ctx, query.OrgId)
	return urList, err
}

func (s *aclService) GetUserRoleById(ctx context.Context, query *accesscontrol.GetUserRoleByIdQuery) (*accesscontrol.UserRole, error) {
	ur, err := s.store.GetUserRoleById(ctx, query.UserRoleId)
	return ur, err
}

func (s *aclService) CreateUserRole(ctx context.Context, cmd *accesscontrol.CreateUserRoleCommand) (*accesscontrol.UserRole, error) {
	ur, err := s.store.CreateUserRole(ctx, cmd.OrgId, cmd.UserRole)
	return ur, err
}

func (s *aclService) UpdateUserRoleById(ctx context.Context, cmd *accesscontrol.UpdateUserRoleByIdCommand) error {
	err := s.store.UpdateUserRoleById(ctx, cmd.UserRoleId, cmd.UserRole)
	return err
}

func (s *aclService) DeleteUserRoleById(ctx context.Context, cmd *accesscontrol.DeleteUserRoleByIdCommand) error {
	err := s.store.DeleteUserRoleById(ctx, cmd.OrgId, cmd.UserRoleId)
	return err
}
