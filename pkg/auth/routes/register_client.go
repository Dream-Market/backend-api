package routes

import (
	"net/http"

	"github.com/Dream-Market/backend-api/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterUserRequestBody struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func RegisterUser(ctx *gin.Context, c pb.AuthServiceClient) (err error) {
	var body RegisterUserRequestBody

	if err = ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Email:    body.Email,
		Phone:    body.Phone,
		Name:     body.Name,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(resp.Status), &resp)
	return
}
