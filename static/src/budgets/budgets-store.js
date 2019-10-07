import { writable } from "svelte/store";

const budgets = writable([
  {
    id: 1,
    name: "cheese",
    category: "grocery",
    price: 2.39,
    type: "out",
    date: "03-10-2019"
  },
  {
    id: 2,
    name: "wage",
    category: "wage",
    price: 1994.95,
    type: "in",
    date: "03-10-2019"
  },
  {
    id: 3,
    name: "rent",
    category: "house",
    price: 1000.0,
    type: "out",
    date: "01-9-2019"
  },
  {
    id: 4,
    name: "rent",
    category: "house",
    price: 1000.0,
    type: "out",
    date: "01-10-2019"
  
  },
  {
    id: 5,
    name: "wage",
    category: "wage",
    price: 1994.95,
    type: "in",
    date: "03-9-2019"
  }
]);

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
