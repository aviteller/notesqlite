package models

import (
	"database/sql"
	"fmt"
)

func InitTables() {
	GetDB().Exec("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY autoincrement, title VARCHAR(255), message TEXT)")

	//  id: 1,
	// 	name: "workout1",
	// 	duration: "10",
	// 	workoutType: "arms",
	// 	actionsNo: 2
	GetDB().Exec("CREATE TABLE IF NOT EXISTS workouts (id INTEGER PRIMARY KEY autoincrement, name text, duration int DEFAULT 0, workout_type text, actions_no int DEFAULT 0, pos int DEFAULT 0)")
	// id:1,
	// workoutID: 1,
	// name: "deadlift",
	// equipment: "barbel",
	// actionType: "time",
	// actionLength: "60"
	GetDB().Exec("CREATE TABLE IF NOT EXISTS actions (id INTEGER PRIMARY KEY autoincrement, workout_id int, name text, equipment TEXT, action_type text, action_length int DEFAULT 0, pos int DEFAULT 0)")
	// id: 3,
	// name: "rent",
	// category: "house",
	// price: 1000.0,
	// type: "out",
	// date: "01-9-2019"
	GetDB().Exec("CREATE TABLE IF NOT EXISTS budgets (id INTEGER PRIMARY KEY autoincrement, name text, category TEXT, price FLOAT, type int DEFAULT 0, date DATE)")
}

//GetDB opens connection to sqlite db
func GetDB() *sql.DB {
	//:databaselocked.sqlite?cache=shared&mode=rwc
	database, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		fmt.Println(err)
	}

	return database
}
