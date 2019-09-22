package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
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
