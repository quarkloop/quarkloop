package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *UserApi) getUserById(ctx *gin.Context, query *user.GetUserByIdQuery) api.Response {
	user, err := s.userService.GetUserById(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, user)
}
