package organization_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/user/store"
)

type userService struct {
	store store.OrgStore
}

func NewOrganizationService(ds store.OrgStore) user.Service {
	return &userService{
		store: ds,
	}
}

func (s *userService) GetUserById(ctx context.Context, query *user.GetUserByIdQuery) (*user.User, error) {
	u, err := s.store.GetUserById(ctx, query.UserId)
	return u, err
}

func (s *userService) GetUserByEmail(ctx context.Context, query *user.GetUserByEmailQuery) (*user.User, error) {
	u, err := s.store.GetUserByEmail(ctx, query.Email)
	return u, err
}

func (s *userService) GetUserAccountByUserId(ctx context.Context, query *user.GetUserAccountByUserIdQuery) (*user.UserAccount, error) {
	u, err := s.store.GetUserAccountByUserId(ctx, query.UserId)
	return u, err
}

func (s *userService) GetUserSessionByUserId(ctx context.Context, query *user.GetUserSessionByUserIdQuery) (*user.UserSession, error) {
	u, err := s.store.GetUserSessionByUserId(ctx, query.UserId)
	return u, err
}

func (s *userService) UpdateUserById(ctx context.Context, cmd *user.UpdateUserByIdCommand) error {
	err := s.store.UpdateUserById(ctx, cmd.UserId, &cmd.User)
	return err
}

func (s *userService) DeleteUserById(ctx context.Context, cmd *user.DeleteUserByIdCommand) error {
	err := s.store.DeleteUserById(ctx, cmd.UserId)
	return err
}
