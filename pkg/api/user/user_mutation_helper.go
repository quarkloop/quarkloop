package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *userApi) updateUser(ctx *gin.Context, cmd *user.UpdateUserCommand) api.Response {
	err := s.userService.UpdateUser(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updateUsername(ctx *gin.Context, cmd *user.UpdateUsernameCommand) api.Response {
	err := s.userService.UpdateUsername(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updatePassword(ctx *gin.Context, cmd *user.UpdatePasswordCommand) api.Response {
	err := s.userService.UpdatePassword(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updatePreferences(ctx *gin.Context, cmd *user.UpdatePreferencesCommand) api.Response {
	err := s.userService.UpdatePreferences(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updateUserById(ctx *gin.Context, cmd *user.UpdateUserByIdCommand) api.Response {
	err := s.userService.UpdateUserById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updateUsernameByUserId(ctx *gin.Context, cmd *user.UpdateUsernameByUserIdCommand) api.Response {
	err := s.userService.UpdateUsernameByUserId(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updatePasswordByUserId(ctx *gin.Context, cmd *user.UpdatePasswordByUserIdCommand) api.Response {
	err := s.userService.UpdatePasswordByUserId(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *userApi) updatePreferencesByUserId(ctx *gin.Context, cmd *user.UpdatePreferencesByUserIdCommand) api.Response {
	err := s.userService.UpdatePreferencesByUserId(ctx, cmd)
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

func (s *userApi) deleteSessionById(ctx *gin.Context, cmd *user.DeleteSessionByIdCommand) api.Response {
	err := s.userService.DeleteSessionById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *userApi) deleteAccountById(ctx *gin.Context, cmd *user.DeleteAccountByIdCommand) api.Response {
	err := s.userService.DeleteAccountById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}
