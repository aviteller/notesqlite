package models

import (
	"database/sql"
	"fmt"

	u "../utils"
)

// id: 3,
// name: "rent",
// category: "house",
// price: 1000.0,
// type: "out",
// date: "01-9-2019"
type Budget struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Type     int     `json:"type"`
	Date     string  `json:"date"`
}

func (budget *Budget) Validate() (map[string]interface{}, bool) {

	if budget.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func (b *Budget) Create() map[string]interface{} {
	if res, ok := b.Validate(); !ok {
		return res
	}
	database := GetDB()
	statement, err := database.Prepare("INSERT INTO `budgets` (`name`,`category`,`price`,`type`,`date`) VALUES (?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(b.Name, b.Category, b.Price, b.Type, b.Date)
	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	b.ID = int(lastid)
	res["budget"] = b
	return res
}

func GetBudgets(month string) []Budget {
	var budgets []Budget
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM budgets WHERE strftime('%m', date) = ?", month)
	for rows.Next() {
		var budget Budget
		_ = rows.Scan(&budget.ID, &budget.Name, &budget.Category, &budget.Price, &budget.Type, &budget.Date)
		budgets = append(budgets, budget)
	}
	rows.Close() //good habit to close
	return budgets
}

func GetBudgetStatsByMonth(month string) []sql.NullString {
	var totalIn sql.NullString
	var totalOut sql.NullString

	database := GetDB()

	err := database.QueryRow("SELECT SUM(`price`) AS total FROM budgets WHERE type = 0 AND strftime('%m', date) = ?", month).Scan(&totalIn)
	if err != nil {
		fmt.Println(err)
	}
	err = database.QueryRow("SELECT SUM(`price`) AS total FROM budgets WHERE type = 1 AND strftime('%m', date) = ?", month).Scan(&totalOut)
	if err != nil {
		fmt.Println(err)
	}

	return []sql.NullString{totalIn, totalOut}
}

func DeleteBudget(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("delete from budgets where id=?")

	stmt.Exec(id)

	res2 := u.Message(true, "success")

	return res2
}

func (b *Budget) UpdateBudget(id string) map[string]interface{} {
	if res, ok := b.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("UPDATE `budgets` SET `name`= ?, `category`=?, `price`=?, `type`=?, `date`=? WHERE id =?")
	result, _ := statement.Exec(b.Name, b.Category, b.Price, b.Type, b.Date, id)

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	b.ID = int(lastid)
	res["budget"] = b
	return res

}
