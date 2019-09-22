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
	Actions     []Action
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
		_ = rows.Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.WorkoutType, &workout.ActionsNo)
		fmt.Println(workout)
		workouts = append(workouts, workout)
	}

	rows.Close() //good habit to close

	return workouts
}

//GetNotes gets all the Notes
func GetWorkoutDetails(id string) Workout {

	var workout Workout
	database := GetDB()
	err := database.QueryRow("SELECT * FROM workouts WHERE id = ?", id).Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.WorkoutType, &workout.ActionsNo)

	if err != nil {
		panic(err.Error())
	}

	actions := GetActionsForWorkout(id)

	workout.Actions = actions
	// for rows.Next() {
	// 	var workout Workout
	// 	_ = rows.Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.WorkoutType, &workout.ActionsNo)
	// 	fmt.Println(workout)
	// 	workouts = append(workouts, workout)
	// }

	// rows.Close() //good habit to close

	return workout
}

func DeleteWorkout(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("delete from workouts where id=?")

	stmt.Exec(id)

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
