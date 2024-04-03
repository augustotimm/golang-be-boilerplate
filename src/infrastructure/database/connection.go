package database

import (
	config "backend-boilerplate/src/application/config/models/contexts"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(dbConfig config.DatabaseConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("DATABASE OK!")

	return db
}
