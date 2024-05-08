package user

import (
	"context"

	userGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/auth/user"
	"github.com/quarkloop/quarkloop/services/user/store"
)

func (s *userService) GetUserById(ctx context.Context, query *userGrpc.GetUserByIdQuery) (*userGrpc.GetUserByIdReply, error) {
	u, err := s.store.GetUserById(ctx, &store.GetUserByIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetUserByIdReply{User: u.ToProto()}
	return res, err
}

func (s *userService) GetUsernameByUserId(ctx context.Context, query *userGrpc.GetUsernameByUserIdQuery) (*userGrpc.GetUsernameByUserIdReply, error) {
	username, err := s.store.GetUsernameByUserId(ctx, &store.GetUsernameByUserIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetUsernameByUserIdReply{Username: username}
	return res, err
}

func (s *userService) GetEmailByUserId(ctx context.Context, query *userGrpc.GetEmailByUserIdQuery) (*userGrpc.GetEmailByUserIdReply, error) {
	email, err := s.store.GetEmailByUserId(ctx, &store.GetEmailByUserIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetEmailByUserIdReply{Email: email}
	return res, err
}

func (s *userService) GetStatusByUserId(ctx context.Context, query *userGrpc.GetStatusByUserIdQuery) (*userGrpc.GetStatusByUserIdReply, error) {
	status, err := s.store.GetStatusByUserId(ctx, &store.GetStatusByUserIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetStatusByUserIdReply{Status: int32(status)}
	return res, err
}

func (s *userService) GetSessionsByUserId(ctx context.Context, query *userGrpc.GetSessionsByUserIdQuery) (*userGrpc.GetSessionsByUserIdReply, error) {
	_, err := s.store.GetSessionsByUserId(ctx, &store.GetSessionsByUserIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetSessionsByUserIdReply{}
	return res, err
}

func (s *userService) GetAccountsByUserId(ctx context.Context, query *userGrpc.GetAccountsByUserIdQuery) (*userGrpc.GetAccountsByUserIdReply, error) {
	_, err := s.store.GetAccountsByUserId(ctx, &store.GetAccountsByUserIdQuery{UserId: query.UserId})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetAccountsByUserIdReply{}
	return res, err
}

func (s *userService) GetUsers(ctx context.Context, query *userGrpc.GetUsersQuery) (*userGrpc.GetUsersReply, error) {
	_, err := s.store.GetUsers(ctx, &store.GetUsersQuery{})
	if err != nil {
		return nil, err
	}

	res := &userGrpc.GetUsersReply{}
	return res, err
}
