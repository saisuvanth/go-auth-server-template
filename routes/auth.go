package routes

import (
	"github.com/gin-gonic/gin"
	"server.com/server/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth_router := router.Group("/auth")
	{
		auth_router.POST("/login", controllers.UserLogin())
		auth_router.POST("/register", controllers.UserRegister())
	}
}
