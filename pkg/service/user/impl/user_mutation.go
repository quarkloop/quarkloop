package user_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/user/store"
)

type userService struct {
	store store.OrgStore
}

func NewUserService(ds store.OrgStore) user.Service {
	return &userService{
		store: ds,
	}
}

func (s *userService) UpdateUserById(ctx context.Context, cmd *user.UpdateUserByIdCommand) error {
	err := s.store.UpdateUserById(ctx, cmd.UserId, &cmd.User)
	return err
}
