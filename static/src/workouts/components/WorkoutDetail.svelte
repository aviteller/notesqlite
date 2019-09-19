<script>
  import { onDestroy, createEventDispatcher } from "svelte";
  import { pop } from "svelte-spa-router";
  import Button from "../../UI/Button.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import ActionGrid from "../../workout-actions/ActionGrid.svelte";
  import workouts from "../workouts-store.js";
  import { isEmpty, isValidEmail } from "../../helpers/validation";
  export let params;

  let selectedWorkout;
  let name = "";
  let duration = "";
  let workoutType = "";
  let actionsNo = "";

  const unsubscribe = workouts.subscribe(items => {
    selectedWorkout = items.find(i => i.id === parseInt(params.id));
    name = selectedWorkout.name;
    duration = selectedWorkout.duration;
    workoutType = selectedWorkout.workoutType;
    actionsNo = selectedWorkout.actionsNo;
  });

  onDestroy(() => unsubscribe());

  const dispatch = createEventDispatcher();

  $: nameValid = !isEmpty(name);
</script>

<style>
  section {
    margin-top: 4rem;
  }
  .content {
    width: 100%;
    margin: auto;
    display: grid;
    grid-template-columns: 20% 77%;
    grid-gap: 3%;
  }

  .edit-form {
    grid-column-start: 1;
    grid-column-end: 2;
  }

  h1 {
    font-size: 2rem;
    font-family: "Roboto Slab", sans-serif;
    margin: 0.5rem 0;
  }

  .action-controls {
    display: flex;
    color: #6b6b6b;
    justify-content: space-around;
  }

  p {
    font-size: 1.5rem;
  }
</style>

<section>

  <div class="content">
    <div class="edit-form">
      <h1>{name}</h1>
      <TextInput
        id="name"
        label="Name"
        value={name}
        valid={nameValid}
        validityMessage="Please enter valid name"
        on:input={e => (name = e.target.value)} />
      <TextInput
        id="duration"
        label="Durtaion"
        value={duration}
        on:input={e => (duration = e.target.value)} />
      <TextInput
        id="workoutType"
        label="Workout Type"
        value={workoutType}
        on:input={e => (workoutType = e.target.value)} />
    </div>
    <div class="actions">
   
        <ActionGrid workoutID={params.id} />
      {#if actionsNo > 0}
        <!-- {#each actions as action} {action.name}{/each} -->
      {:else}
        <h1>Add new action</h1>
      {/if}
    </div>
    <!-- <h2>{selectedMeetup.subtitle} - {selectedMeetup.address}</h2>
    <p>{selectedMeetup.description}</p>
    <Button href="mailto:{selectedMeetup.contactEmail}">Contact</Button> -->
    <Button mode="outline" on:click={() => pop()}>Go Back</Button>
  </div>
</section>
