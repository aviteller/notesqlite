<script>
  import WorkoutItem from "./WorkoutItem.svelte";
  import Button from "../../UI/Button.svelte";
  import { createEventDispatcher, onMount } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  import SortableList from "../../UI/SortableList.svelte";
  export let workouts;

  const dispatch = createEventDispatcher();

  const sortList = ev => {
    workouts = ev.detail;
  };

</script>

<style>
  #meetups {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr;
    grid-gap: 1rem;
  }

  #meetup-controls {
    margin: 1rem;
    display: flex;
    justify-content: space-between;
  }

  #no-meetups {
    margin: 1rem;
  }

  @media (min-width: 768px) {
    #meetups {
      grid-template-columns: repeat(4, 1fr);
    }
  }
</style>

<section id="meetup-controls">
  <Button on:click={() => dispatch('add')}>New Meetup</Button>

</section>
{#if workouts.length === 0}
  <p id="no-meetups">No meetups found</p>
{/if}

<section id="meetups">
  <SortableList list={workouts} key="id" on:sort={sortList} let:item>
    <WorkoutItem
      id={item.id}
      name={item.name}
      duration={item.duration}
      workout_type={item.workout_type}
      actions_no={item.actions_no}
      on:edit
      on:showdetails />
  </SortableList>
  <!-- {#each workouts as workout (workout.id)}
    <div transition:scale animate:flip={{ duration: 300 }}>
      <WorkoutItem
        id={workout.id}
        name={workout.name}
        duration={workout.duration}
        workout_type={workout.workout_type}
        actions_no={workout.actions_no}
        on:edit
        on:showdetails />
    </div>
  {/each} -->
</section>
