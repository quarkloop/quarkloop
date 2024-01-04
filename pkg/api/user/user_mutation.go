package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

// PUT /user
//
// Update current user.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdateUser(ctx *gin.Context) {
	cmd := &user.UpdateUserCommand{}
	res := s.updateUser(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /user/username
//
// Update current user username.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdateUsername(ctx *gin.Context) {
	cmd := &user.UpdateUsernameCommand{}
	res := s.updateUsername(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /user/password
//
// Update current user password.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdatePassword(ctx *gin.Context) {
	cmd := &user.UpdatePasswordCommand{}
	res := s.updatePassword(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /user/preferences
//
// Update current user preferences.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdatePreferences(ctx *gin.Context) {
	cmd := &user.UpdatePreferencesCommand{}
	res := s.updatePreferences(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /users/:userId
//
// Update user by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdateUserById(ctx *gin.Context) {
	uriParams := &user.UpdateUserByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	cmd := &user.UpdateUserByIdCommand{UserId: uriParams.UserId}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res := s.updateUserById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /users/:userId/username
//
// Update user username by user id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdateUsernameByUserId(ctx *gin.Context) {
	uriParams := &user.UpdateUsernameByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	cmd := &user.UpdateUsernameByUserIdCommand{UserId: uriParams.UserId}
	res := s.updateUsernameByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /users/:userId/password
//
// Update user password by user id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdatePasswordByUserId(ctx *gin.Context) {
	uriParams := &user.UpdatePasswordByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	cmd := &user.UpdatePasswordByUserIdCommand{UserId: uriParams.UserId}
	res := s.updatePasswordByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /users/:userId/preferences
//
// Update user preferences by user id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *userApi) UpdatePreferencesByUserId(ctx *gin.Context) {
	uriParams := &user.UpdatePreferencesByUserIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	cmd := &user.UpdatePreferencesByUserIdCommand{UserId: uriParams.UserId}
	res := s.updatePreferencesByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /users/:userId
//
// Delete user by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *userApi) DeleteUserById(ctx *gin.Context) {
	uriParams := &user.DeleteUserByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &user.DeleteUserByIdCommand{UserId: uriParams.UserId}
	res := s.deleteUserById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /users/:userId/sessions/:sessionId
//
// Delete user session by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *userApi) DeleteSessionById(ctx *gin.Context) {
	uriParams := &user.DeleteSessionByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &user.DeleteSessionByIdCommand{
		UserId:    uriParams.UserId,
		SessionId: uriParams.SessionId,
	}
	res := s.deleteSessionById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /users/:userId/accounts/:accountId
//
// Delete user account by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *userApi) DeleteAccountById(ctx *gin.Context) {
	uriParams := &user.DeleteAccountByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &user.DeleteAccountByIdCommand{
		UserId:    uriParams.UserId,
		AccountId: uriParams.AccountId,
	}
	res := s.deleteAccountById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
