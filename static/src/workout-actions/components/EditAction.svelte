<script>
  import { createEventDispatcher } from "svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import Modal from "../../UI/Modal.svelte";
  import { isEmpty, isValidEmail } from "../../helpers/validation";
  import actions from "../actions-store.js";
  import Select from "../../UI/Select.svelte";

  export let id = null;
  export let workoutID;

  let name = "";
  let action_type = "";
  let action_length = "";
  let equipment = "";

  if (id) {
    const unsubscribe = actions.subscribe(items => {
      const selectedAction = items.find(i => i.id === id);
      if (selectedAction) {
        name = selectedAction.name;
        action_type = selectedAction.action_type;
        action_length = selectedAction.action_length.toString();
        equipment = selectedAction.equipment;
      }
    });

    unsubscribe();
  }

  const actionTypes = [
    { value: "time", label: "Time" },
    { value: "reps", label: "Reps" }
  ];
  const handleSelect = e => {
    action_type = e.detail;
  };

  $: nameValid = !isEmpty(name);
  // $: action_typeValid = !isEmpty(action_type);
  $: action_lengthValid = !isEmpty(action_length);
  $: formIsValid = nameValid && action_lengthValid;

  const dispatch = createEventDispatcher();

  const submitForm = () => {
    const newAction = {
      workout_id: parseInt(workoutID),
      name,
      equipment,
      action_type,
      action_length: parseInt(action_length)
    };
    if (id) {
      fetch(`http://localhost:9000/api/actions/${id}`, {
        method: "PUT",
        body: JSON.stringify(newAction),
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
          actions.updateAction(id, newAction);
          cancel();
        })
        .catch(err => console.log(err));
    } else {
      fetch("http://localhost:9000/api/actions", {
        method: "POST",
        body: JSON.stringify(newAction),
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
          actions.addAction({
            ...newAction,
            id: data.action.id
          });
          dispatch("add");
          cancel();
        })
        .catch(err => console.log(err));
    }
  };
  const deleteAction = () => {
    fetch(`http://localhost:9000/api/actions/${id}`, {
      method: "DELETE"
    })
      .then(res => {
        actions.removeAction(id);
        dispatch("remove");
        cancel();
      })
      .catch(err => console.log(err));
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
    <Select
      options={actionTypes}
      selected={action_type}
      on:selectchange={handleSelect}
      label="Type" />
    <TextInput
      id="action_length"
      label="Action Length"
      value={action_length}
      valid={action_lengthValid}
      type="number"
      validityMessage="Please enter valid action_length"
      on:input={e => (action_length = e.target.value)} />
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
