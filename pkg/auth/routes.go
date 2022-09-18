package auth

import (
	"github.com/Dream-Market/backend-api/config"
	"github.com/Dream-Market/backend-api/pkg/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c config.Config) *ServiceClient {
	srvc := ServiceClient{
		Client: InitServiceClient(c),
	}

	router := r.Group("/auth")

	router.POST("/register", srvc.RegisterUser)
	router.POST("/login", srvc.LoginUser)

	return &srvc
}

func (srvc *ServiceClient) RegisterUser(ctx *gin.Context) {
	routes.RegisterUser(ctx, srvc.Client)
}

func (srvc *ServiceClient) LoginUser(ctx *gin.Context) {
	routes.LoginUser(ctx, srvc.Client)
}
