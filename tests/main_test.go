package tests

import (
	dbCleanUp "backend-boilerplate/tests/db"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"os"
	"testing"
)

var db *sql.DB
var mockDB sqlmock.Sqlmock

func TestMain(m *testing.M) {
	setup()
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	db, mockDB, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dbCleanUp.Run(db)
}

func teardown() {
	dbCleanUp.Run(db)
}
