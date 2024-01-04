package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

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

func (s *userApi) UpdateUser(ctx *gin.Context) {
	cmd := &user.UpdateUserCommand{}
	res := s.updateUser(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdateUsername(ctx *gin.Context) {
	cmd := &user.UpdateUsernameCommand{}
	res := s.updateUsername(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdatePassword(ctx *gin.Context) {
	cmd := &user.UpdatePasswordCommand{}
	res := s.updatePassword(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdatePreferences(ctx *gin.Context) {
	cmd := &user.UpdatePreferencesCommand{}
	res := s.updatePreferences(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdateUserById(ctx *gin.Context) {
	cmd := &user.UpdateUserByIdCommand{}
	res := s.updateUserById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdateUsernameByUserId(ctx *gin.Context) {
	cmd := &user.UpdateUsernameByUserIdCommand{}
	res := s.updateUsernameByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdatePasswordByUserId(ctx *gin.Context) {
	cmd := &user.UpdatePasswordByUserIdCommand{}
	res := s.updatePasswordByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) UpdatePreferencesByUserId(ctx *gin.Context) {
	cmd := &user.UpdatePreferencesByUserIdCommand{}
	res := s.updatePreferencesByUserId(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) DeleteUserById(ctx *gin.Context) {
	cmd := &user.DeleteUserByIdCommand{}
	res := s.deleteUserById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) DeleteSessionById(ctx *gin.Context) {
	cmd := &user.DeleteSessionByIdCommand{}
	res := s.deleteSessionById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

func (s *userApi) DeleteAccountById(ctx *gin.Context) {
	cmd := &user.DeleteAccountByIdCommand{}
	res := s.deleteAccountById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
