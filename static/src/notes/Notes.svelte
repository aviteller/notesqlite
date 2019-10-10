<script>
  import NoteTable from "./components/NoteTable.svelte";
  import Button from "../UI/Button.svelte";
  import Modal from "../UI/Modal.svelte";
  import Form from "./components/Form.svelte";
  import notes from "./notes-store.js";

  let editedId = null;
  let editMode = null;
  let isLoading = false;
  const getNotes = () => {
    fetch(`http://localhost:9000/api/notes`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching meetups");
        }
        return res.json();
      })
      .then(data => {
        if (!data.status) {
          error = data;
        } else {
          const loadedNotes = [];
          for (const key in data.data) {
            loadedNotes.push({
              ...data.data[key]
            });
          }
          notes.setNotes(loadedNotes);
          isLoading = false;
        }
      })
      .catch(err => {
        error = err;
        isLoading = false;
      });
  };
  getNotes();

  const startEdit = e => {
    cancelEdit();
    setTimeout(() => {
      
      editMode = "edit";
      editedId = e.detail;
    }, 100);
  };

  const cancelEdit = () => {
    editMode = null;
  };
</script>
<Modal>


<div style="float:right; margin-right:8%;">
  <Button on:click={() => (editMode = 'new')} color="success">New Note</Button>
</div>
{#if editMode == 'edit'}
  <Form ID={editedId} on:cancel={cancelEdit} />
{:else if editMode == 'new'}
  <Form on:cancel={cancelEdit} />
{:else}
  <NoteTable notes={$notes} on:edit={startEdit} />
{/if}
</Modal>

