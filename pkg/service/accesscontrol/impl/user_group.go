package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) GetUserGroupList(ctx context.Context, query *accesscontrol.GetUserGroupListQuery) ([]accesscontrol.UserGroup, error) {
	ugList, err := s.store.GetUserGroupList(ctx, query)
	return ugList, err
}

func (s *aclService) GetUserGroupById(ctx context.Context, query *accesscontrol.GetUserGroupByIdQuery) (*accesscontrol.UserGroup, error) {
	ug, err := s.store.GetUserGroupById(ctx, query)
	return ug, err
}

func (s *aclService) CreateUserGroup(ctx context.Context, cmd *accesscontrol.CreateUserGroupCommand) (*accesscontrol.UserGroup, error) {
	ug, err := s.store.CreateUserGroup(ctx, cmd)
	return ug, err
}

func (s *aclService) UpdateUserGroupById(ctx context.Context, cmd *accesscontrol.UpdateUserGroupByIdCommand) error {
	err := s.store.UpdateUserGroupById(ctx, cmd)
	return err
}

func (s *aclService) DeleteUserGroupById(ctx context.Context, cmd *accesscontrol.DeleteUserGroupByIdCommand) error {
	err := s.store.DeleteUserGroupById(ctx, cmd)
	return err
}
