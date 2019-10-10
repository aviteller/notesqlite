<script>
  import notes from "../notes-store.js";
  import { createEventDispatcher } from "svelte";
  import Modal from "../../UI/Modal.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import { location, pop } from "svelte-spa-router";
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

  const cancelNote = () => {
    if ($location === "/addnote") {
      pop();
    } else {
      dispatch("cancel");
    }
  };

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
            //console.log($location)
            if ($location === "/addnote") {
              pop();
            } else {
              dispatch("cancel");
            }
          })
          .catch(err => console.log(err));
      }
    }
  };
</script>

<style>

</style>

<Modal title={ID ? `Edit ${title}` : 'Add new Note'} on:cancel>

  <TextInput
    id="title"
    label="Title"
    value={title}
    validityMessage="Please fill in correctly"
    on:input={e => (title = e.target.value)} />

  <TextInput
    id="message"
    label="Message"
    value={message}
    type="textarea"
    rows="3"
    bind:value={message} />

  <div slot="footer">
    <Button on:click={onSubmit}>Submit</Button>
    <Button color="danger" on:click={cancelNote}>Cancel</Button>
  </div>

</Modal>
