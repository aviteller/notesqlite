package models

import (
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
	fmt.Println(b)
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

func GetBudgets() []Budget {
	var budgets []Budget
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM budgets")
	for rows.Next() {
		var budget Budget
		_ = rows.Scan(&budget.ID, &budget.Name, &budget.Category, &budget.Price, &budget.Type, &budget.Date)
		budgets = append(budgets, budget)
	}

	rows.Close() //good habit to close

	return budgets
}
