package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mddsilva/gobet/pkg/config"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		user := api.Group("/auth")
		user.POST("/signup", init.AuthCtrl.Signup)
	}

	return router
}
