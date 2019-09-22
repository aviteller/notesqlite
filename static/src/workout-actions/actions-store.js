import { writable } from "svelte/store";

const actions = writable([
  {
    id: 1,
    workoutID: 1,
    name: "deadlift",
    equipment: "barbel",
    actionType: "time",
    actionLength: "60"
  },
  {
    id: 2,
    workoutID: 1,
    name: "sqauts",
    equipment: "barbel",
    actionType: "time",
    actionLength: "30"
  },

  {
    id: 3,
    workoutID: 1,
    name: "deadlift",
    equipment: "barbel",
    actionType: "reps",
    actionLength: "60"
  },
  {
    id: 4,
    workoutID: 2,
    name: "sqauts",
    equipment: "barbel",
    actionType: "time",
    actionLength: "30"
  },
  {
    id: 5,
    workoutID: 2,
    name: "slaps",
    equipment: "barbel",
    actionType: "time",
    actionLength: "30"
  }
]);

const customActionsStore = {
  subscribe: actions.subscribe,
  setActions: (actionArray) => {
    actions.set(actionArray);
  },
  addAction: (actionData) => {
    const newAction = {
      ...actionData
    };
    actions.update((items) => [...items, newAction]);
  },
  updateAction: (id, actionData) => {
    actions.update((items) => {
      const actionIndex = items.findIndex((m) => m.id === id);
      const updatedAction = { ...items[actionIndex], ...actionData };
      const updatedActions = [...items];
      updatedActions[actionIndex] = updatedAction;
      return updatedActions;
    });
  },
  removeAction: (id) => {
    actions.update((items) => {
      return items.filter((i) => i.id !== id);
    });
  }
};

export default customActionsStore;
