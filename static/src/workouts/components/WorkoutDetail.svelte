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
  let workout_type = "";
  let actions_no = "";
  let actions = [];

  fetch(`http://localhost:9000/api/workouts/${parseInt(params.id)}`)
    .then(res => {
      if (!res.ok) {
        throw new Error("Issue fetching workous");
      }
      return res.json();
    })
    .then(data => {
      name = data.data.name;
      duration = data.data.duration;
      workout_type = data.data.workout_type;
      actions_no = data.data.actions_no;
      actions=data.data.Actions
    })
    .catch(err => {
      error = err;
      isLoading = false;
      console.log(err);
    });
  // const unsubscribe = workouts.subscribe(items => {
  //   selectedWorkout = items.find(i => i.id === parseInt(params.id));
  //   name = selectedWorkout.name;
  //   duration = selectedWorkout.duration;
  //   workout_type = selectedWorkout.workout_type;
  //   actions_no = selectedWorkout.actions_no;
  // });

  //   unsubscribe();

  // onDestroy(() => unsubscribe());

  const submitForm = () => {
    const newWorkout = {
      name,
      duration: parseInt(duration),
      workout_type,
      actions_no
    };

    fetch(`http://localhost:9000/api/workouts/${parseInt(params.id)}`, {
      method: "PUT",
      body: JSON.stringify(newWorkout),
      headers: { "Content-Type": "application/json" }
    })
      .then(res => {
        if (!res.ok) {
          throw new Error("Failed");
        }
        return res.json();
      })
      .then(data => {
        if (!data.status) {
          throw new Error(data.message);
        }
        workouts.updateWorkout(parseInt(params.id), newWorkout);
      })
      .catch(err => console.log(err));
  };

  const deleteWorkout = () => {
    workouts.removeWorkout(parseInt(params.id));
    pop();
  };

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

  .workout-buttons {
    display: flex;
    justify-content: space-between;
  }
</style>

<section>

  <div class="content">
    <div class="edit-form">
      <Button mode="outline" on:click={() => pop()}>Go Back</Button>
      <h1>{name}</h1>
      <h2>NO of actions: {actions_no}</h2>
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
        type="number"
        on:input={e => (duration = e.target.value)} />
      <TextInput
        id="workout_type"
        label="Workout Type"
        value={workout_type}
        on:input={e => (workout_type = e.target.value)} />
      <div class="workout-buttons">
        <Button mode="outline" color="success" on:click={submitForm}>
          Update
        </Button>
        <Button mode="outline" color="danger" on:click={deleteWorkout}>
          Delete
        </Button>
      </div>
    </div>
    <div class="actions">

      <ActionGrid actions={actions} workoutID={params.id}/>
      {#if actions_no > 0}
        <!-- {#each actions as action} {action.name}{/each} -->
      {:else}
        <h1>Add new action</h1>
      {/if}
    </div>
    <!-- <h2>{selectedMeetup.subtitle} - {selectedMeetup.address}</h2>
    <p>{selectedMeetup.description}</p>
    <Button href="mailto:{selectedMeetup.contactEmail}">Contact</Button> -->

  </div>
</section>
