package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *userApi) getUser(ctx *gin.Context, query *user.GetUserQuery) api.Response {
	user, err := s.userService.GetUser(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, user)
}

func (s *userApi) getUsername(ctx *gin.Context, query *user.GetUsernameQuery) api.Response {
	username, err := s.userService.GetUsername(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, username)
}

func (s *userApi) getEmail(ctx *gin.Context, query *user.GetEmailQuery) api.Response {
	email, err := s.userService.GetEmail(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, email)
}

func (s *userApi) getStatus(ctx *gin.Context, query *user.GetStatusQuery) api.Response {
	status, err := s.userService.GetStatus(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, status)
}

func (s *userApi) getPreferences(ctx *gin.Context, query *user.GetPreferencesQuery) api.Response {
	preferences, err := s.userService.GetPreferences(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, preferences)
}

func (s *userApi) getSessions(ctx *gin.Context, query *user.GetSessionsQuery) api.Response {
	sessions, err := s.userService.GetSessions(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, sessions)
}

func (s *userApi) getAccounts(ctx *gin.Context, query *user.GetAccountsQuery) api.Response {
	accounts, err := s.userService.GetAccounts(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, accounts)
}

func (s *userApi) getUserById(ctx *gin.Context, query *user.GetUserByIdQuery) api.Response {
	user, err := s.userService.GetUserById(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, user)
}

func (s *userApi) getUsernameByUserId(ctx *gin.Context, query *user.GetUsernameByUserIdQuery) api.Response {
	username, err := s.userService.GetUsernameByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, username)
}

func (s *userApi) getEmailByUserId(ctx *gin.Context, query *user.GetEmailByUserIdQuery) api.Response {
	email, err := s.userService.GetEmailByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, email)
}

func (s *userApi) getStatusByUserId(ctx *gin.Context, query *user.GetStatusByUserIdQuery) api.Response {
	status, err := s.userService.GetStatusByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, status)
}

func (s *userApi) getPreferencesByUserId(ctx *gin.Context, query *user.GetPreferencesByUserIdQuery) api.Response {
	preferences, err := s.userService.GetPreferencesByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, preferences)
}

func (s *userApi) getSessionsByUserId(ctx *gin.Context, query *user.GetSessionsByUserIdQuery) api.Response {
	sessions, err := s.userService.GetSessionsByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, sessions)
}

func (s *userApi) getAccountsByUserId(ctx *gin.Context, query *user.GetAccountsByUserIdQuery) api.Response {
	accounts, err := s.userService.GetAccountsByUserId(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, accounts)
}

func (s *userApi) getUsers(ctx *gin.Context, query *user.GetUsersQuery) api.Response {
	users, err := s.userService.GetUsers(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, users)
}
