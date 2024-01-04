package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *userApi) GetUser(ctx *gin.Context) {
	query := &user.GetUserQuery{}
	res := s.getUser(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetUsername(ctx *gin.Context) {
	query := &user.GetUsernameQuery{}
	res := s.getUsername(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetEmail(ctx *gin.Context) {
	query := &user.GetEmailQuery{}
	res := s.getEmail(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetStatus(ctx *gin.Context) {
	query := &user.GetStatusQuery{}
	res := s.getStatus(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetPreferences(ctx *gin.Context) {
	query := &user.GetPreferencesQuery{}
	res := s.getPreferences(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetSessions(ctx *gin.Context) {
	query := &user.GetSessionsQuery{}
	res := s.getSessions(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetAccounts(ctx *gin.Context) {
	query := &user.GetAccountsQuery{}
	res := s.getAccounts(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /users/:userId
//
// Get user by id.
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

func (s *userApi) GetUsernameByUserId(ctx *gin.Context) {
	query := &user.GetUsernameByUserIdQuery{}
	res := s.getUsernameByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetEmailByUserId(ctx *gin.Context) {
	query := &user.GetEmailByUserIdQuery{}
	res := s.getEmailByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetStatusByUserId(ctx *gin.Context) {
	query := &user.GetStatusByUserIdQuery{}
	res := s.getStatusByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetPreferencesByUserId(ctx *gin.Context) {
	query := &user.GetPreferencesByUserIdQuery{}
	res := s.getPreferencesByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetSessionsByUserId(ctx *gin.Context) {
	query := &user.GetSessionsByUserIdQuery{}
	res := s.getSessionsByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetAccountsByUserId(ctx *gin.Context) {
	query := &user.GetAccountsByUserIdQuery{}
	res := s.getAccountsByUserId(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) GetUsers(ctx *gin.Context) {
	query := &user.GetUsersQuery{}
	res := s.getUsers(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
