package models

import (
	u "../utils"
	//blank
	_ "github.com/mattn/go-sqlite3"
)

//Note struct exported
type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (note *Note) Validate() (map[string]interface{}, bool) {

	if note.Title == "" || note.Message == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func (note *Note) Create() map[string]interface{} {
	if res, ok := note.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `notes` (`title`,`message`) VALUES (?,?)")
	result, _ := statement.Exec(note.Title, note.Message)
	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	note.ID = int(lastid)
	res["note"] = note
	return res
}

//GetNotes gets all the Notes
func GetNotes() []Note {
	var notes []Note
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM notes")

	for rows.Next() {
		var note Note
		_ = rows.Scan(&note.ID, &note.Title, &note.Message)
		notes = append(notes, note)
	}

	rows.Close() //good habit to close

	return notes
}

func DeleteNote(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("delete from notes where id=?")

	stmt.Exec(id)

	res2 := u.Message(true, "success")

	return res2
}

func (note *Note) UpdateNote(id string) map[string]interface{} {
	if res, ok := note.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("UPDATE `notes` SET `title`= ?, `message`=? WHERE id =?")
	result, _ := statement.Exec(note.Title, note.Message, id)

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	note.ID = int(lastid)
	res["note"] = note
	return res

}
