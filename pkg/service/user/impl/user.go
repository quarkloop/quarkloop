package user_impl

import (
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/user/store"
)

type userService struct {
	store store.UserStore
}

func NewUserService(ds store.UserStore) user.Service {
	return &userService{
		store: ds,
	}
}
