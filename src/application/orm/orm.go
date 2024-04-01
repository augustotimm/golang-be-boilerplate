package orm

import (
	"database/sql"
	"fmt"
	"go-be-boilerplate/src/application/config"
)

func ConnectDatabase() (*sql.DB, error) {
	config := config.Config()

	connectString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name)

	db, err := sql.Open(config.Database.Name, connectString)
	if err != nil {
		fmt.Printf("Fail connecting to database: %s", err.Error())
	}

	return db, err
}
