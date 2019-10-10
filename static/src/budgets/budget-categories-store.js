import { writable } from "svelte/store";

const budgetCategories = writable([]);

const customBudgetCategoriesStore = {
  subscribe: budgetCategories.subscribe,
  setBudgetCategories: budgetCategoriesArray => {

    budgetCategoriesArray.forEach(category => {
      category["label"] = category["name"]
      delete  category["name"]
      category["value"] = category["id"]
      delete  category["id"]
      delete category["type"]
      delete category["deleted"]
    });

    budgetCategories.set(budgetCategoriesArray);
  },
  addBudgetCategories: budgetData => {
    const newBudgetCategories = {
      ...budgetData
    };
    budgetCategories.update(items => [...items, newBudgetCategories]);
  },
  updateBudgetCategories: (id, budgetData) => {
    budgetCategories.update(items => {
      const budgetIndex = items.findIndex(m => m.id === id);
      const updatedBudgetCategory = { ...items[budgetIndex], ...budgetData };
      const updatedBudgetCategories = [...items];
      updatedBudgetCategories[budgetIndex] = updatedBudgetCategory;
      return updatedBudgetCategories;
    });
  },
  removeBudgetCategories: id => {
    budgetCategories.update(items => {
      return items.filter(i => i.id !== id);
    });
  }
};

export default customBudgetCategoriesStore;
