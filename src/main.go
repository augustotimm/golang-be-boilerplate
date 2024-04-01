package main

import (
	"go-be-boilerplate/src/application/api"
	"go-be-boilerplate/src/application/config"
	"go-be-boilerplate/src/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Config()

	db := database.Connect(config.Database)

	ginApp := gin.Default()
	api.Setup(ginApp, db)
	api.Server(config.Api.Port, ginApp)
}
