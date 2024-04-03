package tests

import (
	"database/sql"
)

func Run(db *sql.DB) {
	db.Exec("DELETE FROM hello_world_entities")
}
