package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) ListUserRoles(ctx context.Context, orgId int) ([]accesscontrol.UserRole, error) {
	urList, err := s.store.ListUserRoles(ctx, orgId)
	return urList, err
}

func (s *aclService) GetUserRoleById(ctx context.Context, userRoleId int) (*accesscontrol.UserRole, error) {
	ur, err := s.store.GetUserRoleById(ctx, userRoleId)
	return ur, err
}

func (s *aclService) CreateUserRole(ctx context.Context, orgId int, userRole *accesscontrol.UserRole) (*accesscontrol.UserRole, error) {
	ur, err := s.store.CreateUserRole(ctx, orgId, userRole)
	return ur, err
}

func (s *aclService) UpdateUserRoleById(ctx context.Context, userRoleId int, userRole *accesscontrol.UserRole) error {
	err := s.store.UpdateUserRoleById(ctx, userRoleId, userRole)
	return err
}

func (s *aclService) DeleteUserRoleById(ctx context.Context, orgId int, userRoleId int) error {
	err := s.store.DeleteUserRoleById(ctx, orgId, userRoleId)
	return err
}
