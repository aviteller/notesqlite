package models

import (
	"fmt"

	u "../utils"
)

type Action struct {
	ID           int    `json:"id"`
	WorkoutID    int    `json:"workout_id"`
	Name         string `json:"name"`
	Equipment    string `json:"equipment"`
	ActionType   string `json:"action_type"`
	ActionLength int    `json:"action_length"`
	Pos          int    `json:"pos"`
}

func (action *Action) Validate() (map[string]interface{}, bool) {

	if action.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func GetActionsForWorkout(id string) []Action {
	var actions []Action
	database := GetDB()

	rows, _ := database.Query("SELECT * FROM actions WHERE workout_id = ?", id)

	for rows.Next() {
		var action Action
		err := rows.Scan(&action.ID, &action.WorkoutID, &action.Name, &action.Equipment, &action.ActionType, &action.ActionLength, &action.Pos)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(action)
		actions = append(actions, action)
	}
	rows.Close() //good habit to close
	return actions
}

func (action *Action) Create() map[string]interface{} {
	if res, ok := action.Validate(); !ok {
		return res
	}
	fmt.Println(action.WorkoutID, action.Name, action.Equipment, action.ActionType, action.ActionLength)
	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `actions` (`workout_id`, `name`,`equipment`,`action_type`,`action_length`) VALUES (?,?,?,?,?)")
	result, err := statement.Exec(action.WorkoutID, action.Name, action.Equipment, action.ActionType, action.ActionLength)

	UpdateActionNos(action.WorkoutID, "+1")

	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	action.ID = int(lastid)
	res["action"] = action
	return res
}

func getWorkID(id string) int {
	var workoutId int
	database := GetDB()
	err := database.QueryRow("SELECT workout_id FROM actions WHERE id = ?", id).Scan(&workoutId)

	if err != nil {
		panic(err.Error())
	}

	//workoutIdInt, _ := strconv.Atoi(workoutId)

	return workoutId
}

func DeleteAction(id string) map[string]interface{} {

	workoutID := getWorkID(id)
	UpdateActionNos(workoutID, "-1")
	database := GetDB()
	stmt, _ := database.Prepare("delete from actions where id=?")

	stmt.Exec(id)
	res2 := u.Message(true, "success")

	return res2
}

func (action *Action) UpdateAction(id string) map[string]interface{} {
	if res, ok := action.Validate(); !ok {
		return res
	}
	database := GetDB()
	statement, err := database.Prepare("UPDATE `actions` SET `name`= ?, `equipment`= ?, `action_type`=?, `action_length`=? WHERE id =?")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(action.Name, action.Equipment, action.ActionType, action.ActionLength, id)
	if err != nil {
		fmt.Println(err)
	}

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	action.ID = int(lastid)
	res["action"] = action
	return res

}
