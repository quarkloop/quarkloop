package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

// GET /users/:userId
//
// Get user by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *UserApi) GetUserById(ctx *gin.Context) {
	uriParams := &user.GetUserByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	query := &user.GetUserByIdQuery{UserId: uriParams.UserId}
	res := s.getUserById(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
