package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
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
	params := mux.Vars(r)

	month := params["month"]

	data := models.GetBudgets(month)

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
var GetBudgetStatsByMonth = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	month := params["month"]

	data := models.GetBudgetStatsByMonth(month)

	res := u.Message(true, "success")
	if len(data) > 0 {
		for i, stat := range data {
			if i == 0 {
				res["in"] = stat.String
			} else {
				res["out"] = stat.String
			}
		}
	}

	u.Respond(w, res)

}

var UpdateBudget = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	note := &models.Budget{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := note.UpdateBudget(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteBudget = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteBudget(id)
	res["data"] = id
	u.Respond(w, res)
}
