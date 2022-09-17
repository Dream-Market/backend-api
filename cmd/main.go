package main

import (
	"log"

	"github.com/Dream-Market/backend-api/config"
	"github.com/Dream-Market/backend-api/pkg/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, c)

	r.Run(c.Port)
}
