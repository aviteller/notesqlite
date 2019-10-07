<script>
  import Button from "../../UI/Button.svelte";
  import Badge from "../../UI/Badge.svelte";
  import LoadingSpinner from "../../UI/LoadingSpinner.svelte";
  import { createEventDispatcher } from "svelte";

  import workouts from "../workouts-store.js";
  export let name;
  export let id;
  export let duration;
  export let workout_type;
  export let actions_no;

  let isLoading = false;

  const dispatch = createEventDispatcher();

  const exportWorkout = () => {
    fetch(`http://localhost:9000/api/workoutexport/${id}`)
      .then(res => {
       
        window.open(res.url, '_blank')
        
      })
      .catch(err => {
        console.log(err);
      });
  };
</script>

<style>
  article {
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
    border-radius: 5px;
    background: white;
    margin: 1rem;
  }

  header,
  .content,
  footer {
    padding: 1rem;
    text-align: center;
  }

  h1 {
    font-size: 1.25rem;
    margin: 0.5rem 0;
    font-family: "Roboto Slab", sans-serif;
  }

  /* h1.is-favorite {
    background: #01a129;
    color: white;
    padding: 0 0.5rem;
    border-radius: 5px;
  } */

  h2 {
    font-size: 1rem;
    color: #808080;
    margin: 0.5rem 0;
  }

  p {
    font-size: 1.25rem;
    margin: 0;
  }

  div {
    text-align: right;
  }

  .content {
    height: 4rem;
  }
</style>

<article>
  <header>
    <h1>{name}</h1>
    {#if duration}
      <h2>Time: {duration}</h2>
    {/if}
    {#if workout_type}
      <p>Workout type: {workout_type}</p>
    {/if}
  </header>

  <div class="content">
    {#if actions_no}
      <p>No of actions: {actions_no}</p>
    {/if}
  </div>
  <footer>

    {#if isLoading}
      <span>Changing...</span>
    {:else}
      <!-- <Button
        mode="outline"
        color={isLiked ? 'null' : 'success'}
        on:click={toggleLike}>
        {isLiked ? 'UnLike' : 'Like'}
      </Button> -->
    {/if}
    <Button href={`/workouts/${id}`}>Show Details</Button>
    <Button on:click={exportWorkout}>Print</Button>
    <!-- <Button on:click={() => dispatch('showdetails', id)}>Show Details</Button> -->

  </footer>
</article>
