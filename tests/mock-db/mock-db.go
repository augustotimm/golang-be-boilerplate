package tests

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

var DB *sql.DB
var Mock sqlmock.Sqlmock

func Setup() {
	var err error
	DB, Mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dbCleanUp(DB)
}

func Teardown() {
	dbCleanUp(DB)
}

func dbCleanUp(db *sql.DB) {
	db.Exec("DELETE FROM hello_world_entities")
}
