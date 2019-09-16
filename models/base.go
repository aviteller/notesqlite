package models

import (
	"database/sql"
	"fmt"
)

//GetDB opens connection to sqlite db
func GetDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./dbs/notes.db")
	if err != nil {
		fmt.Println(err)
	}

	return database
}
