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

func (s *userService) GetUserById(ctx context.Context, params *user.GetUserByIdParams) (*user.User, error) {
	u, err := s.store.GetUserById(ctx, params.UserId)
	return u, err
}

func (s *userService) GetUserByEmail(ctx context.Context, params *user.GetUserByEmailParams) (*user.User, error) {
	u, err := s.store.GetUserByEmail(ctx, params.Email)
	return u, err
}

func (s *userService) GetUserAccountByUserId(ctx context.Context, params *user.GetUserAccountByUserIdParams) (*user.UserAccount, error) {
	u, err := s.store.GetUserAccountByUserId(ctx, params.UserId)
	return u, err
}

func (s *userService) GetUserSessionByUserId(ctx context.Context, params *user.GetUserSessionByUserIdParams) (*user.UserSession, error) {
	u, err := s.store.GetUserSessionByUserId(ctx, params.UserId)
	return u, err
}

func (s *userService) UpdateUserById(ctx context.Context, params *user.UpdateUserByIdParams) error {
	err := s.store.UpdateUserById(ctx, params.UserId, &params.User)
	return err
}

func (s *userService) DeleteUserById(ctx context.Context, params *user.DeleteUserByIdParams) error {
	err := s.store.DeleteUserById(ctx, params.UserId)
	return err
}
