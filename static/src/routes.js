import Notes from "./notes/Notes.svelte";
import Workout from "./workouts/Workouts.svelte";
import Budget from "./budgets/Budgets.svelte";
import WorkoutDetail from "./workouts/components/WorkoutDetail.svelte";

// This demonstrates how to pass routes as a POJO (Plain Old JavaScript Object) or a JS Map
// In this code sample we're using both (controlling at runtime what's enabled, by checking for the 'routemap=1' querystring parameter) just because we are using this code sample for tests too
// In your code, you'll likely want to choose one of the two options only
let routes;
const urlParams = new URLSearchParams(window.location.search);
if (!urlParams.has("routemap")) {
  // The simples way to define routes is to use a dictionary.
  // This is a key->value pair in which the key is a string of the path.
  // The path is passed to regexparam that does some transformations allowing the use of params and wildcards
  routes = {
    // Exact path
    "/": Notes,

    // Allow children to also signal link activation
    "/workouts": Workout,

    // Using named parameters, with last being optional
    "/workouts/:id": WorkoutDetail,

    "/budgets": Budget,

    // Using named parameters, with last being optional
  };
} else {
  routes = new Map();

  // Exact path
  routes.set("/", Home);

  // Allow children to also signal link activation
  routes.set("/workouts", Workout);

  // Using named parameters, with last being optional
  routes.set("/workouts/:id", WorkoutDetail);

  // Regular expressions
  routes.set(/^\/regex\/(.*)?/i, Regex);
  routes.set(/^\/(pattern|match)(\/[a-z0-9]+)?/i, Regex);

  // Catch-all, must be last
  // routes.set("*", NotFound);
}

export default routes;
