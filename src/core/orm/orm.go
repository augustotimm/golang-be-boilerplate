package orm

import (
	"backend-boilerplate/src/application/config"
	"backend-boilerplate/src/infrastructure/database"
	"database/sql"
)

func ConnectDatabase() *sql.DB {
	envConfig := config.Config()

	db := database.Connect(envConfig.Database)

	return db
}
