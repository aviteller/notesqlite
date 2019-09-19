<script>
  import ActionItem from "./ActionItem.svelte";
  //import MeetupFilter from "./MeetupFilter.svelte";
  import Button from "../UI/Button.svelte";
  import { onDestroy, createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  import actions from "./actions-store.js";

  const dispatch = createEventDispatcher();

  export let workoutID;

  let actionsArray = [];

  const unsubscribe = actions.subscribe(items => {
    actionsArray = items.filter(i => i.workoutID !== parseInt(workoutID));
  });

  console.log(actionsArray);

  onDestroy(() => unsubscribe());
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

{#if actionsArray.length === 0}
  <p id="no-meetups">No meetups found</p>
{/if}
<section id="meetups">
  {#each actionsArray as action (action.id)}

    <div transition:scale animate:flip={{ duration: 300 }}>
      <ActionItem
        id={action.id}
        name={action.name}
        actionLength={action.actionLength}
        actionType={action.actionType}
        equipment={action.equipment}
        workoutID={action.workoutID}
        on:edit />
    </div>
  {/each}
</section>
