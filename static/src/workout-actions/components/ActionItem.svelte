<script>
  import Button from "../../UI/Button.svelte";
  import { createEventDispatcher } from "svelte";

  import workouts from "../actions-store.js";
  export let name;
  export let id;
  export let action_length;
  export let action_type;
  export let workoutID;
  export let equipment;

  let isLoading = false;

  const dispatch = createEventDispatcher();

</script>

<style>
  article {
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
    border-radius: 5px;
    background: white;
    margin: 1rem;
    padding: 1rem;
  }

  header,
  footer {
    text-align: center;
  }

  h1 {
    font-size: 1.25rem;
    margin: 0.5rem 0;
    font-family: "Roboto Slab", sans-serif;
  }

  h2 {
    font-size: 1rem;
    color: #808080;
    margin: 0.5rem 0;
  }
</style>

<article>
  <header>
    <h1>{name}</h1>
    <h2>{action_type} - {action_length}</h2>
    <h2>{equipment}</h2>
  </header>
  <footer>
    <Button color="success" on:click={() => dispatch('edit', id)}>Edit</Button>
  </footer>
</article>
