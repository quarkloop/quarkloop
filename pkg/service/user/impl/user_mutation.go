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

func (s *userService) UpdateUser(ctx context.Context, cmd *user.UpdateUserCommand) error {
	err := s.store.UpdateUser(ctx, cmd)
	return err
}

func (s *userService) UpdateUsername(ctx context.Context, cmd *user.UpdateUsernameCommand) error {
	err := s.store.UpdateUsername(ctx, cmd)
	return err
}

func (s *userService) UpdatePassword(ctx context.Context, cmd *user.UpdatePasswordCommand) error {
	err := s.store.UpdatePassword(ctx, cmd)
	return err
}

func (s *userService) UpdatePreferences(ctx context.Context, cmd *user.UpdatePreferencesCommand) error {
	err := s.store.UpdatePreferences(ctx, cmd)
	return err
}

func (s *userService) UpdateUserById(ctx context.Context, cmd *user.UpdateUserByIdCommand) error {
	err := s.store.UpdateUserById(ctx, cmd)
	return err
}

func (s *userService) UpdateUsernameByUserId(ctx context.Context, cmd *user.UpdateUsernameByUserIdCommand) error {
	err := s.store.UpdateUsernameByUserId(ctx, cmd)
	return err
}

func (s *userService) UpdatePasswordByUserId(ctx context.Context, cmd *user.UpdatePasswordByUserIdCommand) error {
	err := s.store.UpdatePasswordByUserId(ctx, cmd)
	return err
}

func (s *userService) UpdatePreferencesByUserId(ctx context.Context, cmd *user.UpdatePreferencesByUserIdCommand) error {
	err := s.store.UpdatePreferencesByUserId(ctx, cmd)
	return err
}

func (s *userService) DeleteUserById(ctx context.Context, cmd *user.DeleteUserByIdCommand) error {
	err := s.store.DeleteUserById(ctx, cmd)
	return err
}

func (s *userService) DeleteSessionById(ctx context.Context, cmd *user.DeleteSessionByIdCommand) error {
	err := s.store.DeleteSessionById(ctx, cmd)
	return err
}

func (s *userService) DeleteAccountById(ctx context.Context, cmd *user.DeleteAccountByIdCommand) error {
	err := s.store.DeleteAccountById(ctx, cmd)
	return err
}
