package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"server.com/server/configs"
	"server.com/server/routes"
)

func main() {

	app := gin.Default()

	go configs.ConnectDB()

	defer func() {
		if err := configs.DB.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	router := app.Group("/")

	routes.Init(router)

	app.Run(":8080")
}
