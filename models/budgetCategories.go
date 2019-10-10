package models

import (
	"fmt"

	u "../utils"
)

type BudgetCategory struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Deleted bool   `json:"deleted"`
}

func (bc *BudgetCategory) Validate() (map[string]interface{}, bool) {

	if bc.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func (bc *BudgetCategory) Create() map[string]interface{} {
	if res, ok := bc.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `budget_categories` (`name`,`type`) VALUES (?,?)")
	result, _ := statement.Exec(bc.Name, bc.Type)
	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	bc.ID = int(lastid)
	res["budget_category"] = bc
	return res
}

//GetNotes gets all the Notes
func GetBudgetCategories() []BudgetCategory {
	var bcs []BudgetCategory
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM budget_categories WHERE deleted = 0")

	for rows.Next() {
		var bc BudgetCategory
		_ = rows.Scan(&bc.ID, &bc.Name)
		bcs = append(bcs, bc)
	}

	rows.Close() //good habit to close

	return bcs
}
func GetBudgetCategoriesByType(typeString string) []BudgetCategory {
	var bcs []BudgetCategory
	database := GetDB()
	rows, err := database.Query("SELECT id, name FROM budget_categories WHERE deleted = 0 and `type` = ?", typeString)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var bc BudgetCategory
		_ = rows.Scan(&bc.ID, &bc.Name)
		bcs = append(bcs, bc)
	}

	rows.Close() //good habit to close

	return bcs
}

//GetNotes gets all the Notes
func GetDeletedBudgetCategories() []BudgetCategory {
	var bcs []BudgetCategory
	database := GetDB()
	rows, _ := database.Query("SELECT * FROM budget_categories WHERE deleted = 1")

	for rows.Next() {
		var bc BudgetCategory
		_ = rows.Scan(&bc.ID, &bc.Name)
		bcs = append(bcs, bc)
	}

	rows.Close() //good habit to close

	return bcs
}

func DeleteBudgetCategory(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("UPDATE `budget_categories` SET `deleted`= 1 WHERE id =?")

	stmt.Exec(id)

	res := u.Message(true, "success")

	return res
}

func (bc *BudgetCategory) UpdateBudgetCategory(id string) map[string]interface{} {
	if res, ok := bc.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("UPDATE `budget_categories` SET `name`= ? WHERE id =?")
	result, _ := statement.Exec(bc.Name, id)

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	bc.ID = int(lastid)
	res["budget_category"] = bc
	return res

}
