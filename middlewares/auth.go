package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.com/server/configs"
	"server.com/server/utils"
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a_h := ctx.Request.Header.Get("Authorization")
		if a_h != "" {
			fmt.Println(a_h)
			token := strings.Split(a_h, "Bearer ")[1]
			var user utils.UserResponse
			payload, t_err := jwt.DecodeSegment(token)
			if t_err != nil {
				log.Fatal(t_err)
				ctx.Abort()
				ctx.JSON(http.StatusUnauthorized, utils.AppError{Message: "User not Authorized"})
			}
			user_id, _ := primitive.ObjectIDFromHex(string(payload))
			coll := configs.GetCollection("users")
			u_err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: user_id}}).Decode(&user)
			if u_err != nil {
				log.Fatal(u_err)
				ctx.Abort()
				ctx.JSON(http.StatusUnauthorized, utils.AppError{Message: "User not Authorized"})
			}
			fmt.Println(user)
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, utils.AppError{Message: "User not Authorized"})
		}
	}
}
