import { writable } from "svelte/store";

const workouts = writable([
  {
    id:1,
    name:"workout1",
    duration:"10",
    workoutType:"arms",
    actionsNo:2
  },
  {
    id:2,
    name:"workout2",
    duration:"15",
    workoutType:"arms",
    actionsNo: 3
  }
]);

const customWorkoutsStore = {
  subscribe: workouts.subscribe,
  setWorkouts: workoutArray => {
    workouts.set(workoutArray);
  },
  addWorkout: workoutData => {
    const newWorkout = {
      ...workoutData
    };
    workouts.update(items => [...items, newWorkout]);
  },
  updateWorkout: (id, workoutData) => {
    workouts.update(items => {
      const workoutIndex = items.findIndex(m => m.id === id);
      const updatedWorkout = { ...items[workoutIndex], ...workoutData };
      const updatedWorkouts = [...items];
      updatedWorkouts[workoutIndex] = updatedWorkout;
      return updatedworkouts;
    });
  },
  removeWorkout: id => {
    workouts.update(items => {
      return items.filter(i => i.id !== id);
    });
  }
};

export default customWorkoutsStore;
