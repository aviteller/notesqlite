package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

//CreateNote is an action the creates a note and then sends the data tot eh model
var CreateAction = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	action := &models.Action{}

	err := json.NewDecoder(r.Body).Decode(action)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := action.Create()
	u.Respond(w, res)
}

var GetActionsForWorkout = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	actionid := params["actionid"]

	data := models.GetActionsForWorkout(actionid)

	res := u.Message(true, "success")

	res["data"] = data

	u.Respond(w, res)

}

var UpdateAction = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	action := &models.Action{}

	err := json.NewDecoder(r.Body).Decode(action)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := action.UpdateAction(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteAction = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteAction(id)
	res["data"] = id
	u.Respond(w, res)
}
