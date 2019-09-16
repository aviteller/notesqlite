<script>
  import notes from "../store.js";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let ID = null;

  let title = "";
  let message = "";

  if (ID) {
    const unsubscribe = notes.subscribe(items => {
      const selectedNote = items.find(i => i.id === ID);
      title = selectedNote.title;
      message = selectedNote.message;
    });

    unsubscribe();
  }

  const onSubmit = () => {
    let newNote = {
      title,
      message
    };

    if (!newNote.title || !newNote.message) {
      alert("please fill in all fields");
    } else {
      if (ID) {
        newNote.ID = ID;
        fetch(`http://localhost:9000/api/notes/${ID}`, {
          method: "PUT",
          body: JSON.stringify(newNote),
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
            notes.updateNote(ID, newNote);
            dispatch("cancel");
          })
          .catch(err => console.log(err));
      } else {
        fetch("http://localhost:9000/api/notes", {
          method: "POST",
          body: JSON.stringify(newNote),
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
            notes.addNote({ ...newNote, id: data.note.id });
            title = "";
            message = "";
          })
          .catch(err => console.log(err));
      }
    }
  };
</script>

<style>
  .form-inline {
    display: flex;
    flex-flow: row wrap;
    align-items: center;
    justify-content: space-around;
  }

  .form-inline label {
    margin: 5px 10px 5px 0;
  }

  /* Style the input fields */
  .form-inline input {
    vertical-align: middle;
    margin: 5px 10px 5px 0;
    padding: 10px;
    background-color: #fff;
    border: 1px solid #ddd;
  }

  /* Style the submit button */
  .form-inline button {
    padding: 10px 20px;
    background-color: dodgerblue;
    border: 1px solid #ddd;
    color: white;
  }

  .form-inline button:hover {
    background-color: royalblue;
  }

  @media (max-width: 800px) {
    .form-inline input {
      margin: 10px 0;
    }

    .form-inline {
      flex-direction: column;
      align-items: stretch;
    }
  }
</style>

<div class="form-inline">
  <label for="title">Title</label>
  <input type="text" id="title" bind:value={title} />
  <label for="message">Message</label>
  <textarea id="message" cols="30" rows="3" bind:value={message} />
  <button on:click={onSubmit}>Submit</button>
  {#if ID}
    <button on:click={() => dispatch('cancel')}>Cancel</button>
  {/if}
</div>
