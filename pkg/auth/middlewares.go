package auth

import (
	"net/http"
	"strings"

	"github.com/Dream-Market/backend-api/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	srvc *ServiceClient
}

func InitAuthMiddleware(srvc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{srvc}
}

func (c *AuthMiddlewareConfig) WithAuth(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.srvc.Client.ValidateUser(ctx, &pb.ValidateUserRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("customerId", res.UserId)

	ctx.Next()
	return
}
