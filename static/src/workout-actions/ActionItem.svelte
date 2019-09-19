<script>
  import Button from "../UI/Button.svelte";
  import Badge from "../UI/Badge.svelte";
  import LoadingSpinner from "../UI/LoadingSpinner.svelte";
  import { createEventDispatcher } from "svelte";

  import workouts from "./actions-store.js";
  export let name;
  export let id;
  export let actionLength;
  export let actionType;
  export let workoutID;
  export let equipment;

  let isLoading = false;

  const dispatch = createEventDispatcher();

  // const toggleLike = () => {
  //   isLoading = true;
  //   fetch(`http://localhost:9000/api/meetups/togglelike/${id}`, {
  //     method: "PUT"
  //   })
  //     .then(res => {
  //       isLoading = false;

  //       meetups.toggleLike(id);
  //     })
  //     .catch(err => {
  //       isLoading = false;
  //       console.log(err);
  //     });
  // };
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



  div {
    text-align: right;
  }

  .content {
    height: 4rem;
  }
</style>

<article>
  <header>
    <h1>
      {name}
    </h1>
    <h2>{actionType}  - {actionLength}</h2>
    <h2>{equipment}</h2>
    
    <!-- {#if duration}
    <h2>Time: {duration}</h2>
    {/if}
    {#if workoutType}
    <p>Workout type: {workoutType}</p>
    {/if} -->
  </header>

 
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
    <Button color="success">Edit</Button>
    <!-- <Button on:click={() => dispatch('showdetails', id)}>Show Details</Button> -->

  </footer>
</article>
