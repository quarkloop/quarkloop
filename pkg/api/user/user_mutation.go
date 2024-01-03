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

func (s *UserApi) UpdateUserById(ctx *gin.Context) {
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

func (s *UserApi) DeleteUserById(ctx *gin.Context) {
	uriParams := &user.DeleteUserByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &user.DeleteUserByIdCommand{UserId: uriParams.UserId}
	res := s.deleteUserById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
