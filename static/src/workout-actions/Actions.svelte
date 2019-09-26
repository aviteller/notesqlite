<script>
  import ActionGrid from "./components/ActionGrid.svelte";
  import EditAction from "./components/EditAction.svelte";
  import Button from "../UI/Button.svelte";
  import actions from "./actions-store";

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
    console.log(e.detail)
    editMode = true;
    editedID = e.detail;
  };
  const addAction = () => (editMode = true);
  const stopEdit = () => (editMode = false);
</script>

<style>
  .action-controls {
    display: flex;
    color: #6b6b6b;
    justify-content: space-around;
  }
</style>

<div class="action-controls">
  <h1>Actions</h1>
  <Button on:click={addAction}>Add Action</Button>
</div>
{#if editMode}
  <EditAction id={editedID} {workoutID} on:cancel={stopEdit} on:add on:remove />
{:else if !isLoading}
  <ActionGrid actions={$actions} on:edit={starEdit} />
{/if}
