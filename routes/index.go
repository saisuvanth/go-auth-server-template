package routes

import "github.com/gin-gonic/gin"

func Init(router *gin.RouterGroup) {
	AuthRoutes(router)
	UserRoutes(router)
}
