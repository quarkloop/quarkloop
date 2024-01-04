package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *userApi) updateUserById(ctx *gin.Context, cmd *user.UpdateUserByIdCommand) api.Response {
	err := s.userService.UpdateUserById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) deleteUserById(ctx *gin.Context, cmd *user.DeleteUserByIdCommand) api.Response {
	err := s.userService.DeleteUserById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}
