<script>
  import Header from "./UI/Header.svelte";
  import { Router, Route } from "svelte-routing";
  import Notes from "./notes/Notes.svelte";
  import Workout from "./Workout.svelte";
  export let url = "";
  let title = "notes";
 
</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>

<Router {url}>
  <Header>{title}</Header>
  <main>
    <Route path="/" component={Notes} />
    <Route path="/workouts" component={Workout} />
  </main>
</Router>
