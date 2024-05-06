package user

import (
	"google.golang.org/grpc"

	userGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/auth/user"
	"github.com/quarkloop/quarkloop/services/user/store"
)

type UserService interface {
	userGrpc.UserServiceServer

	RegisterService(s *grpc.Server)
}

type userService struct {
	store store.UserStore

	userGrpc.UnimplementedUserServiceServer
}

func NewUserService(ds store.UserStore) UserService {
	return &userService{store: ds}
}

func (service *userService) RegisterService(s *grpc.Server) {
	userGrpc.RegisterUserServiceServer(s, service)
}
