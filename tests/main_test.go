package tests

import (
	ConfigContexts "go-be-boilerplate/src/application/config/models/contexts"
	"go-be-boilerplate/src/infrastructure/database"
	dbCleanUp "go-be-boilerplate/tests/db"
	"os"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

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
		Name:     "goApiStoreDB",
	}

	db = database.Connect(localConfig)

	dbCleanUp.Run(db)
}

func teardown() {
	dbCleanUp.Run(db)
}
