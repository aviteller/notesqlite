<script>
  import WorkoutItem from "./WorkoutItem.svelte";
  //import MeetupFilter from "./MeetupFilter.svelte";
  import Button from "../../UI/Button.svelte";
  import { createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";

  const dispatch = createEventDispatcher();

  export let workouts;

  let filter = false;

  // $: filteredMeetups = filter
  //   ? filter === 1
  //     ? meetups.filter(m => m.isLiked)
  //     : meetups.filter(m => !m.isLiked)
  //   : meetups;

  const setFilter = e => {
    filter = e.detail;
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
  <!-- <MeetupFilter on:select={setFilter} /> -->
  <Button on:click={() => dispatch('add')}>New Meetup</Button>

</section>
{#if workouts.length === 0}
  <p id="no-meetups">No meetups found</p>
{/if}
<section id="meetups">
  {#each workouts as workout (workout.id)}
    <div transition:scale animate:flip={{ duration: 300 }}>
      <WorkoutItem
        id={workout.id}
        name={workout.name}
        duration={workout.duration}
        workoutType={workout.workoutType}
        actionsNo={workout.actionsNo}
        on:edit
        on:showdetails />
    </div>
  {/each}
</section>
