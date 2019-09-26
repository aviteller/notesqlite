package models

import (
	"fmt"

	u "../utils"
)

type Workout struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Duration    int    `json:"duration"`
	WorkoutType string `json:"workout_type"`
	ActionsNo   int    `json:"actions_no"`
	Pos         int    `json:"pos"`
}

func (workout *Workout) Validate() (map[string]interface{}, bool) {

	if workout.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func (workout *Workout) Create() map[string]interface{} {
	if res, ok := workout.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `workouts` (`name`,`duration`,`workout_type`) VALUES (?,?,?)")
	result, _ := statement.Exec(workout.Name, workout.Duration, workout.WorkoutType)
	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	workout.ID = int(lastid)
	res["workout"] = workout
	return res
}

//GetNotes gets all the Notes
func GetWorkouts() []Workout {
	var workouts []Workout
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM workouts")
	for rows.Next() {
		var workout Workout
		_ = rows.Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.WorkoutType, &workout.ActionsNo, &workout.Pos)
		workouts = append(workouts, workout)
	}

	rows.Close() //good habit to close

	return workouts
}

//GetNotes gets all the Notes
func GetWorkoutDetails(id string) Workout {

	var workout Workout
	database := GetDB()
	err := database.QueryRow("SELECT * FROM workouts WHERE id = ?", id).Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.WorkoutType, &workout.ActionsNo, &workout.Pos)

	if err != nil {
		panic(err.Error())
	}

	return workout
}

func GetActionsNoByWorkoutID(workoutID int) int {
	var actionNo int
	err := GetDB().QueryRow("SELECT `actions_no` FROM workouts WHERE id = ?", workoutID).Scan(&actionNo)

	if err != nil {
		fmt.Println(err)
	}

	return actionNo
}

func UpdateActionNos(id int, change string) {
	s := fmt.Sprintf("UPDATE `workouts` SET `actions_no` = `actions_no` %s   WHERE id = %v", change, id)
	//database := GetDB()
	GetDB().Exec(s)
	// statement, _ := database.Prepare("UPDATE `workouts` SET `actions_no` = `actions_no` ?   WHERE id = ?")
	// statement.Exec(change, id)
}

func DeleteWorkout(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("delete from actions where workout_id=?")
	stmt.Exec(id)
	stmt2, _ := database.Prepare("delete from workouts where id=?")
	stmt2.Exec(id)

	res2 := u.Message(true, "success")

	return res2
}

func (workout *Workout) UpdateWorkout(id string) map[string]interface{} {
	if res, ok := workout.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("UPDATE `workouts` SET `name`= ?, `duration`=? , `workout_type`=?, `actions_no`=? WHERE id =?")
	result, _ := statement.Exec(workout.Name, workout.Duration, workout.WorkoutType, workout.ActionsNo, id)

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	workout.ID = int(lastid)
	res["workout"] = workout
	return res

}
