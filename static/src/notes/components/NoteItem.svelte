<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../../UI/Button.svelte";
  import notes from "../notes-store.js";

  export let ID;
  export let title;
  export let message;
  const dispatch = createEventDispatcher();

  const removeNote = ID => {
    if (window.confirm(`Are you sure you want to delete note: ${title}`)) {
      fetch(`http://localhost:9000/api/notes/${ID}`, {
        method: "DELETE"
      })
        .then(res => {
          notes.removeNote(ID);
        })
        .catch(err => console.log(err));
    }
  };
</script>

<style>
  td {
    text-align: center;
  }
</style>

<td>{ID}</td>
<td>{title}</td>
<td>{message}</td>
<td>
  <Button on:click={() => dispatch('edit', ID)}>EDIT</Button>
  <Button color="danger" on:click={() => removeNote(ID)}>X</Button>
</td>
