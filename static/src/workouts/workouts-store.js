import { writable } from "svelte/store";

const workouts = writable([
  {
    name:"workout1",
    duration:"10",
    workoutType:"arms",
    actions:[
      {
        name:"deadlift",
        equipment:"barbel",
        actionType:"time",
        actionLength:"60",
      },
      {
        name:"sqauts",
        equipment:"barbel",
        actionType:"time",
        actionLength:"30",
      },
    ]
  },
  {
    name:"workout2",
    duration:"15",
    workoutType:"arms",
    actions:[
      {
        name:"deadlift",
        equipment:"barbel",
        actionType:"time",
        actionLength:"60",
      },
      {
        name:"sqauts",
        equipment:"barbel",
        actionType:"time",
        actionLength:"30",
      },
    ]
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
