package main

import (
	"github.com/gin-gonic/gin"
	"go-be-boilerplate/src/application/api"
	"go-be-boilerplate/src/application/config"
	"go-be-boilerplate/src/application/orm"
)

func main() {
	envConfig := config.Config()

	db, err := orm.ConnectDatabase()
	if err != nil {
		panic("error connecting to Database")
	}

	ginApp := gin.Default()
	api.Setup(ginApp, db)
	api.Server(envConfig.Api.Port, ginApp)
}
