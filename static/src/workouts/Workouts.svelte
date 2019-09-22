<script>

  import WorkoutGrid from "./components/WorkoutGrid.svelte";
  import EditWorkout from "./components/EditWorkout.svelte";
  import CreateWorkout from "./components/CreateWorkout.svelte";
  import WorkoutDetail from "./components/WorkoutDetail.svelte";
  import TextInput from "../UI/TextInput.svelte";
  import Button from "../UI/Button.svelte";
  import Error from "../UI/Error.svelte";
  import LoadingSpinner from "../UI/LoadingSpinner.svelte";
  import workouts from "./workouts-store.js";



  //let loadedMeetups = [];

  let editMode = null;
  let editedId = null;
  let page = "overview";
  let pageData = {};
  let isLoading = false;
  let error;

  fetch("http://localhost:9000/api/workouts")
    .then(res => {
      if (!res.ok) {
        throw new Error("Issue fetching workous");
      }
      return res.json();
    })
    .then(data => {
      const loadedWorkouts = [];
      for (const key in data.data) {
        loadedWorkouts.push({
          ...data.data[key],
          // id: key
        });
      }

      isLoading = false;
      workouts.setWorkouts(loadedWorkouts);
    })
    .catch(err => {
      error = err;
      isLoading = false;
      console.log(err);
    });

  const savedWorkout = () => {
    editMode = null;
    editedId = null;
  };
  const cancelEdit = () => {
    editMode = null;
    editedId = null;
  };
  const showDetails = e => {
    page = "details";
    pageData.id = e.detail;
 
  };

  const close = () => {
    page = "overview";
    pageData.id = {};
  };

  const startEdit = e => {
    editMode = "edit";
    editedId = e.detail;
  };

  const clearError = () => {
    error = null;
  }
</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>
{#if error}
<Error message={error.message} on:cancel={clearError}/>
{/if}


<main>
  {#if page === 'overview'}
    {#if editMode === 'edit'}
      <CreateWorkout
        id={editedId}
        on:addworkout={savedWorkout}
        on:cancel={cancelEdit} />
    {/if}
    {#if isLoading}
      <LoadingSpinner />
    {:else}

      <WorkoutGrid
        workouts={$workouts}
        on:showdetails={showDetails}
        on:edit={startEdit}
        on:add={() => (editMode = 'edit')} />
    {/if}
  {:else}
    <WorkoutDetail id={pageData.id} on:close={close} />
  {/if}
</main>
