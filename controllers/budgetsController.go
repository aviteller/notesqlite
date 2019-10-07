package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
)

var CreateBudget = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	budget := &models.Budget{}

	err := json.NewDecoder(r.Body).Decode(budget)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := budget.Create()
	u.Respond(w, res)
}

var GetBudgets = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetBudgets()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
