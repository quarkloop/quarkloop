package contextdata

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func GetUser(ctx context.Context) *user.User {
	user, ok := ctx.Value(userKey).(*user.User)
	if !ok || user == nil {
		panic("user must be available")
	}

	return user
}

func SetUser(ctx context.Context, u *user.User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func IsUserAnonymous(ctx context.Context) bool {
	user, ok := ctx.Value(userKey).(*user.User)
	if !ok || user == nil {
		return true
	}

	return false
}
