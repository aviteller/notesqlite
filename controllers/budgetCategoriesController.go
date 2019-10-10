package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

var CreateBudgetCategory = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	bc := &models.BudgetCategory{}

	err := json.NewDecoder(r.Body).Decode(bc)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := bc.Create()
	u.Respond(w, res)
}

//GetNotes gets all the birthdays
var GetBudgetCategories = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetBudgetCategories()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var GetBudgetCategoriesByType = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	typeString := params["type"]

	data := models.GetBudgetCategoriesByType(typeString)

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var UpdateBudgetCategory = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	bc := &models.BudgetCategory{}

	err2 := json.NewDecoder(r.Body).Decode(bc)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := bc.UpdateBudgetCategory(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteBudgetCategory = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteBudgetCategory(id)
	res["data"] = id
	u.Respond(w, res)
}
