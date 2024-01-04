package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

// GET /user
//
// Get current user.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetUser(ctx *gin.Context) {
	query := &user.GetUserQuery{}
	res := s.getUser(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/username
//
// Get current user username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetUsername(ctx *gin.Context) {
	query := &user.GetUsernameQuery{}
	res := s.getUsername(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/email
//
// Get current user email address.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetEmail(ctx *gin.Context) {
	query := &user.GetEmailQuery{}
	res := s.getEmail(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/status
//
// Get current user status.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetStatus(ctx *gin.Context) {
	query := &user.GetStatusQuery{}
	res := s.getStatus(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/preferences
//
// Get current user preferences.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetPreferences(ctx *gin.Context) {
	query := &user.GetPreferencesQuery{}
	res := s.getPreferences(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/sessions
//
// Get current user sessions.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetSessions(ctx *gin.Context) {
	query := &user.GetSessionsQuery{}
	res := s.getSessions(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /user/accounts
//
// Get current user accounts.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetAccounts(ctx *gin.Context) {
	query := &user.GetAccountsQuery{}
	res := s.getAccounts(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username
//
// Get user by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetUserById(ctx *gin.Context) {
	uriParams := &user.GetUserByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetUserByIdQuery{UserId: uriParams.UserId}
	res := s.getUserById(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/username
//
// Get user username by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetUsernameByUserId(ctx *gin.Context) {
	uriParams := &user.GetUsernameByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetUsernameByUserIdQuery{UserId: uriParams.UserId}
	res := s.getUsernameByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/email
//
// Get user email address by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetEmailByUserId(ctx *gin.Context) {
	uriParams := &user.GetEmailByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetEmailByUserIdQuery{UserId: uriParams.UserId}
	res := s.getEmailByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/status
//
// Get user status by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetStatusByUserId(ctx *gin.Context) {
	uriParams := &user.GetStatusByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetStatusByUserIdQuery{UserId: uriParams.UserId}
	res := s.getStatusByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/preferences
//
// Get user preferences by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetPreferencesByUserId(ctx *gin.Context) {
	uriParams := &user.GetPreferencesByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetPreferencesByUserIdQuery{UserId: uriParams.UserId}
	res := s.getPreferencesByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/sessions
//
// Get user sessions by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetSessionsByUserId(ctx *gin.Context) {
	uriParams := &user.GetSessionsByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetSessionsByUserIdQuery{UserId: uriParams.UserId}
	res := s.getSessionsByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId_or_username/accounts
//
// Get user accounts by id or username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetAccountsByUserId(ctx *gin.Context) {
	uriParams := &user.GetAccountsByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetAccountsByUserIdQuery{UserId: uriParams.UserId}
	res := s.getAccountsByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users
//
// Get users.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) GetUsers(ctx *gin.Context) {
	query := &user.GetUsersQuery{}
	res := s.getUsers(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
