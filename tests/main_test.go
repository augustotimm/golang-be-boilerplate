package tests

import (
	ConfigContexts "backend-boilerplate/src/application/config/models/contexts"
	"backend-boilerplate/src/infrastructure/database"
	dbCleanUp "backend-boilerplate/tests/db"
	"database/sql"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	localConfig := ConfigContexts.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "root",
		Name:     "postgres",
	}

	db = database.Connect(localConfig)

	dbCleanUp.Run(db)
}

func teardown() {
	dbCleanUp.Run(db)
}
