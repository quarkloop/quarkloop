package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) ListUserGroups(ctx context.Context, orgId int) ([]accesscontrol.UserGroup, error) {
	ugList, err := s.store.ListUserGroups(ctx, orgId)
	return ugList, err
}

func (s *aclService) GetUserGroupById(ctx context.Context, userGroupId int) (*accesscontrol.UserGroup, error) {
	ug, err := s.store.GetUserGroupById(ctx, userGroupId)
	return ug, err
}

func (s *aclService) CreateUserGroup(ctx context.Context, orgId int, userGroup *accesscontrol.UserGroup) (*accesscontrol.UserGroup, error) {
	ug, err := s.store.CreateUserGroup(ctx, orgId, userGroup)
	return ug, err
}

func (s *aclService) UpdateUserGroupById(ctx context.Context, userGroupId int, userGroup *accesscontrol.UserGroup) error {
	err := s.store.UpdateUserGroupById(ctx, userGroupId, userGroup)
	return err
}

func (s *aclService) DeleteUserGroupById(ctx context.Context, orgId int, userGroupId int) error {
	err := s.store.DeleteUserGroupById(ctx, orgId, userGroupId)
	return err
}
