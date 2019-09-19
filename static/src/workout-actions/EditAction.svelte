<script>
  import { createEventDispatcher } from "svelte";
  import TextInput from "../UI/TextInput.svelte";
  import Button from "../UI/Button.svelte";
  import Modal from "../UI/Modal.svelte";
  import { isEmpty, isValidEmail } from "../helpers/validation";
  import actions from "./actions-store.js";

  export let id = null;
  export let workoutID = null;

  let name = "";
  let actionType = "";
  let actionLength = "";
  let equipment = "";


  if (id) {
    const unsubscribe = actions.subscribe(items => {
      const selectedAction = items.find(i => i.id === id);
      if(selectedAction){

        name = selectedAction.name;
        actionType = selectedAction.actionType;
        actionLength = selectedAction.actionLength;
        equipment = selectedAction.equipment;
      }
    });

    unsubscribe();
  }

  $: nameValid = !isEmpty(name);
  $: actionTypeValid = !isEmpty(actionType);
  $: actionLengthValid = !isEmpty(actionLength);
  $: formIsValid = nameValid && actionTypeValid && actionLengthValid;

  const dispatch = createEventDispatcher();

  //   const submitForm = e => {

  //   editMode = null;
  // };

  const submitForm = () => {
    const newAction = {
      name,
      actionType,
      actionLength,
      equipment
    };
    if (id) {
      // fetch(`http://localhost:9000/api/meetups/${id}`, {
      //   method: "PUT",
      //   body: JSON.stringify(newMeetup),
      //   headers: { "Content-Type": "application/json" }
      // })
      //   .then(res => {
      //     if (!res.ok) {
      //       throw new Error("Failed");
      //     }
      //     return res.json();
      //   })
      //   .then(data => {
      //     if (!data.status) {
      //       throw new Error(data.message);
      //     }
      //     meetups.updateMeetup(id, newMeetup);
      //     dispatch("addmeetup");
      //     clearForm();
      //   })
      //   .catch(err => console.log(err));
      actions.updateAction(id, newAction);
      cancel();
    } else {
      // fetch("http://localhost:9000/api/meetups", {
      //   method: "POST",
      //   body: JSON.stringify(newMeetup),
      //   headers: { "Content-Type": "application/json" }
      // })
      //   .then(res => {
      //     if (!res.ok) {
      //       throw new Error("Failed");
      //     }
      //     return res.json();
      //   })
      //   .then(data => {
      //     if (!data.status) {
      //       throw new Error(data.message);
      //     }
      //     meetups.addMeetup({
      //       ...newMeetup,
      //       isLiked: false,
      //       id: data.meetup.ID
      //     });
      //     dispatch("addmeetup");

      //   })
      //   .catch(err => console.log(err));
      actions.addAction({
        ...newAction,
        id: Math.random()
      });
      cancel();
    }
  };
  const deleteAction = () => {
    actions.removeAction(id);
    cancel();
    // fetch(`http://localhost:9000/api/meetups/${id}`, {
    //   method: "DELETE"
    // })
    //   .then(res => {
    //     meetups.removeMeetup(id);
    //     cancel();
    //   })
    //   .catch(err => console.log(err));
  };
  const cancel = () => dispatch("cancel");
</script>

<style>
  form {
    max-width: 100%;
  }
</style>

<Modal title={id ? `Edit ${name}` : 'New Action'} on:cancel>
  <form>
    <TextInput
      id="name"
      label="Name"
      value={name}
      valid={nameValid}
      validityMessage="Please enter valid name"
      on:input={e => (name = e.target.value)} />
    <TextInput
      id="actionType"
      label="Type"
      value={actionType}
      valid={actionTypeValid}
      validityMessage="Please enter valid actionType"
      on:input={e => (actionType = e.target.value)} />
    <TextInput
      id="actionLength"
      label="Address"
      value={actionLength}
      valid={actionLengthValid}
      validityMessage="Please enter valid actionLength"
      on:input={e => (actionLength = e.target.value)} />

    <TextInput
      id="equipment"
      label="Equipment"
      value={equipment}
      on:input={e => (equipment = e.target.value)} />
  </form>
  <div slot="footer">
    <Button on:click={cancel}>Cancel</Button>
    <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
    {#if id}
      <Button on:click={deleteAction}>Delete</Button>
    {/if}
  </div>
</Modal>
