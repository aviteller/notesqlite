<script>
  import ActionItem from "./ActionItem.svelte";
  import EditAction from "./EditAction.svelte";
  //import MeetupFilter from "./MeetupFilter.svelte";
  import Button from "../UI/Button.svelte";
  import { onDestroy, createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  import actions from "./actions-store";

  const dispatch = createEventDispatcher();

  export let workoutID;
  let isLoading = true;
  let editMode = false;
  let editedID = null;

  fetch(`http://localhost:9000/api/workouts/actions/${workoutID}`)
    .then(res => {
      if (!res.ok) {
        throw new Error("Issue fetching workous");
      }
      return res.json();
    })
    .then(data => {
      const loadedActions = [];
      for (const key in data.data) {
        loadedActions.push({
          ...data.data[key]
          // id: key
        });
      }

      isLoading = false;
      actions.setActions(loadedActions);
    })
    .catch(err => {
      isLoading = false;
      console.log(err);
    });

  const starEdit = e => {
    editMode = true;
    editedID = e.detail;
  };

  const addAction = () => (editMode = true);
  const stopEdit = () => (editMode = false);
</script>

<style>
  #meetups {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr;
    grid-gap: 1rem;
  }

  @media (min-width: 768px) {
    #meetups {
      grid-template-columns: repeat(4, 1fr);
    }
  }

  .action-controls {
    display: flex;
    color: #6b6b6b;
    justify-content: space-around;
  }
</style>

{#if editMode}
  <EditAction id={editedID} {workoutID} on:cancel={stopEdit} on:add on:remove/>
{:else}
  {#if isLoading}
    <h1>...LOADING</h1>
  {:else}
    <div class="action-controls">
      <h1>Actions</h1>
      <Button on:click={addAction}>Add Action</Button>
    </div>
    <section id="meetups">
      {#if actions && $actions.length > 0}
        {#each $actions as action (action.id)}
          <div transition:scale animate:flip={{ duration: 300 }}>
            <ActionItem
              id={action.id}
              name={action.name}
              action_length={action.action_length}
              action_type={action.action_type}
              equipment={action.equipment}
              workoutID={action.workoutID}
              on:edit={starEdit} />
          </div>
        {/each}
      {:else}
        <h1>Add actions</h1>
      {/if}
    </section>
  {/if}
{/if}
