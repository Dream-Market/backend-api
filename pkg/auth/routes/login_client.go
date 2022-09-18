package routes

import (
	"net/http"

	"github.com/Dream-Market/backend-api/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginUserRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginUser(ctx *gin.Context, c pb.AuthServiceClient) (err error) {
	var body LoginUserRequestBody

	if err = ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.LoginUser(ctx, &pb.LoginUserRequest{
		Login:    body.Login,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(resp.Status), &resp)
	return
}
