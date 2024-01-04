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

func (s *userService) GetUserById(ctx context.Context, query *user.GetUserByIdQuery) (*user.User, error) {
	u, err := s.store.GetUserById(ctx, query.UserId)
	return u, err
}

/////////////

func (s *userService) GetUser(ctx context.Context, query *user.GetUserQuery) (*user.User, error) {
	res, err := s.store.GetUser(ctx, query)
	return res, err
}

func (s *userService) GetUsername(ctx context.Context, query *user.GetUsernameQuery) (string, error) {
	res, err := s.store.GetUsername(ctx, query)
	return res, err
}

func (s *userService) GetEmail(ctx context.Context, query *user.GetEmailQuery) (string, error) {
	res, err := s.store.GetEmail(ctx, query)
	return res, err
}

func (s *userService) GetStatus(ctx context.Context, query *user.GetStatusQuery) (any, error) {
	res, err := s.store.GetStatus(ctx, query)
	return res, err
}

func (s *userService) GetPreferences(ctx context.Context, query *user.GetPreferencesQuery) (any, error) {
	res, err := s.store.GetPreferences(ctx, query)
	return res, err
}

func (s *userService) GetSessions(ctx context.Context, query *user.GetSessionsQuery) ([]*user.UserSession, error) {
	res, err := s.store.GetSessions(ctx, query)
	return res, err
}

func (s *userService) GetAccounts(ctx context.Context, query *user.GetAccountsQuery) ([]*user.UserAccount, error) {
	res, err := s.store.GetAccounts(ctx, query)
	return res, err
}

func (s *userService) GetUserById(ctx context.Context, query *user.GetUserByIdQuery) (*user.User, error) {
	res, err := s.store.GetUserById(ctx, query)
	return res, err
}

func (s *userService) GetUsernameByUserId(ctx context.Context, query *user.GetUsernameByUserIdQuery) (string, error) {
	res, err := s.store.GetUsernameByUserId(ctx, query)
	return res, err
}

func (s *userService) GetEmailByUserId(ctx context.Context, query *user.GetEmailByUserIdQuery) (string, error) {
	res, err := s.store.GetEmailByUserId(ctx, query)
	return res, err
}

func (s *userService) GetStatusByUserId(ctx context.Context, query *user.GetStatusByUserIdQuery) (any, error) {
	res, err := s.store.GetStatusByUserId(ctx, query)
	return res, err
}

func (s *userService) GetPreferencesByUserId(ctx context.Context, query *user.GetPreferencesByUserIdQuery) (any, error) {
	res, err := s.store.GetPreferencesByUserId(ctx, query)
	return res, err
}

func (s *userService) GetSessionsByUserId(ctx context.Context, query *user.GetSessionsByUserIdQuery) ([]*user.UserSession, error) {
	res, err := s.store.GetSessionsByUserId(ctx, query)
	return res, err
}

func (s *userService) GetAccountsByUserId(ctx context.Context, query *user.GetAccountsByUserIdQuery) ([]*user.UserAccount, error) {
	res, err := s.store.GetAccountsByUserId(ctx, query)
	return res, err
}

func (s *userService) GetUsers(ctx context.Context, query *user.GetUsersQuery) ([]*user.User, error) {
	res, err := s.store.GetUsers(ctx, query)
	return res, err
}
