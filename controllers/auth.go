package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server.com/server/configs"
	"server.com/server/models"
	"server.com/server/utils"
)

func UserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tok_chan := make(chan string)
		var form utils.LoginForm
		if ctx.ShouldBind(&form) == nil {
			var user models.User
			coll := configs.GetCollection("users")
			filter := bson.D{{Key: "username", Value: form.Username}}
			err := coll.FindOne(context.TODO(), filter).Decode(&user)
			go utils.GenToken(tok_chan, user.Id)
			token := <-tok_chan
			if err != nil {
				log.Fatal(err)
				ctx.JSON(400, "No user found")
			}
			if utils.ComparePassword(form.Password, user.Password) {
				ctx.JSON(200, token)
			} else {
				ctx.JSON(401, map[string]any{"message": "User credentials doesnt match"})
			}
		}
	}
}

func UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pass_chan := make(chan string)
		var form utils.RegisterForm
		if ctx.ShouldBind(&form) == nil {
			coll := configs.GetCollection("users")
			user := models.User{Username: form.Username, Password: form.Password, Email: form.Email}
			go utils.EncryptPassword(pass_chan, user.Password)
			res, err := coll.InsertOne(context.TODO(), user)
			if err != nil {
				log.Fatal(err)
				ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
			}
			hash_password := <-pass_chan
			update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hash_password}}}}
			_, u_err := coll.UpdateByID(context.TODO(), res.InsertedID, update)
			if u_err != nil {
				log.Fatal(u_err)
			}
			ctx.JSON(200, map[string]any{"user_id": res.InsertedID})
		}
	}
}
