package routes

import (
	"github.com/gin-gonic/gin"
	"server.com/server/controllers"
	"server.com/server/middlewares"
)

func UserRoutes(router *gin.RouterGroup) {
	user_router := router.Group("/", middlewares.AuthCheck())
	{
		user_router.GET("/", controllers.GetUser())
	}
}
