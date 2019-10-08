import { writable } from "svelte/store";

const budgets = writable();

const customBudgetsStore = {
  subscribe: budgets.subscribe,
  setBudgets: budgetArray => {
    budgets.set(budgetArray);
  },
  addBudget: budgetData => {
    const newBudget = {
      ...budgetData
    };
    budgets.update(items => [...items, newBudget]);
  },
  updateBudget: (id, budgetData) => {
    budgets.update(items => {
      const budgetIndex = items.findIndex(m => m.id === id);
      const updatedBudget = { ...items[budgetIndex], ...budgetData };
      const updatedBudgets = [...items];
      updatedBudgets[budgetIndex] = updatedBudget;
      return updatedBudgets;
    });
  },
  removeBudget: id => {
    budgets.update(items => {
      return items.filter(i => i.id !== id);
    });
  }
};

export default customBudgetsStore;
