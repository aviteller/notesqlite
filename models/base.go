package models

import (
	"database/sql"
	"fmt"
)

func InitTables() {
	GetDB().Exec("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY autoincrement, title VARCHAR(255), message TEXT)")

}

//GetDB opens connection to sqlite db
func GetDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./dbs/notes.db")
	if err != nil {
		fmt.Println(err)
	}

	return database
}
