package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")

		ctx.JSON(http.StatusOK, user)
	}
}
