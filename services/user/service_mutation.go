package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	userGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/auth/user"
	"github.com/quarkloop/quarkloop/services/user/store"
)

func (s *userService) UpdateUserById(ctx context.Context, cmd *userGrpc.UpdateUserByIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateUserById(ctx, &store.UpdateUserByIdCommand{})
	return &emptypb.Empty{}, err
}

func (s *userService) UpdateUsernameByUserId(ctx context.Context, cmd *userGrpc.UpdateUsernameByUserIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateUsernameByUserId(ctx, &store.UpdateUsernameByUserIdCommand{})
	return &emptypb.Empty{}, err
}

func (s *userService) UpdatePasswordByUserId(ctx context.Context, cmd *userGrpc.UpdatePasswordByUserIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdatePasswordByUserId(ctx, &store.UpdatePasswordByUserIdCommand{})
	return &emptypb.Empty{}, err
}

func (s *userService) DeleteUserById(ctx context.Context, cmd *userGrpc.DeleteUserByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteUserById(ctx, &store.DeleteUserByIdCommand{})
	return &emptypb.Empty{}, err
}

func (s *userService) DeleteSessionById(ctx context.Context, cmd *userGrpc.DeleteSessionByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteSessionById(ctx, &store.DeleteSessionByIdCommand{})
	return &emptypb.Empty{}, err
}

func (s *userService) DeleteAccountById(ctx context.Context, cmd *userGrpc.DeleteAccountByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteAccountById(ctx, &store.DeleteAccountByIdCommand{})
	return &emptypb.Empty{}, err
}
