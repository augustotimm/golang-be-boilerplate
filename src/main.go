package main

import (
	"backend-boilerplate/src/application/api"
	"backend-boilerplate/src/application/config"
	"backend-boilerplate/src/core/orm"
	"github.com/gin-gonic/gin"
)

func main() {
	envConfig := config.Config()

	db := orm.ConnectDatabase()

	ginApp := gin.Default()
	api.Setup(ginApp, db)
	api.Server(envConfig.Api.Port, ginApp)
}
