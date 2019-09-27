package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

//CreateNote is an action the creates a note and then sends the data tot eh model
var CreateWorkout = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	workout := &models.Workout{}

	err := json.NewDecoder(r.Body).Decode(workout)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := workout.Create()
	u.Respond(w, res)
}

//GetNotes gets all the birthdays
var GetWorkouts = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetWorkouts()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var GetWorkoutDetails = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	data := models.GetWorkoutDetails(id)

	res := u.Message(true, "success")

	res["data"] = data

	u.Respond(w, res)

}

var UpdateWorkout = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	workout := &models.Workout{}

	err2 := json.NewDecoder(r.Body).Decode(workout)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := workout.UpdateWorkout(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteWorkout = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteWorkout(id)
	res["data"] = id
	u.Respond(w, res)
}
var ExportWorkout = func(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data := models.ExportWorkout(id)
	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename=workout-"+id+".csv")
	w.Write(data.Bytes())
}
