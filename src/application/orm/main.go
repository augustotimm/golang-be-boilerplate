package main

import (
	"go-be-boilerplate/src/application/config"
	"go-be-boilerplate/src/core/entities"
	"go-be-boilerplate/src/infrastructure/database"
)

func main() {
	config := config.Config()

	db := database.Connect(config.Database)

	db.AutoMigrate(&entities.HelloWorldEntity{})
}
