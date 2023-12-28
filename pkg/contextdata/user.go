package contextdata

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func GetUser(ctx *gin.Context) *user.User {
	val, exists := ctx.Get(userKey)
	if !exists || val == nil {
		panic("user must be available")
	}

	u, ok := val.(*user.User)
	if !ok {
		panic("*user.User type assertion failed")
	}

	return u
}

func SetUser(ctx *gin.Context, u *user.User) {
	ctx.Set(userKey, u)
}

func IsUserAnonymous(ctx *gin.Context) bool {
	val, exists := ctx.Get(userKey)
	if !exists || val == nil {
		return true
	}

	_, ok := val.(*user.User)
	return !ok
}
