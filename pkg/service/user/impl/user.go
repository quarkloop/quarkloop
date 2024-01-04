package user_impl

import (
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
