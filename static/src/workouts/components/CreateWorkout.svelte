<script>
  import { createEventDispatcher } from "svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import Modal from "../../UI/Modal.svelte";
  import { isEmpty, isValidEmail } from "../../helpers/validation";
  import workouts from "../workouts-store.js";
  let name = "";
  const dispatch = createEventDispatcher();
    const cancel = () => dispatch("cancel");

  $: nameValid = !isEmpty(name);
  $: formIsValid = nameValid;

  const submitForm = () => {
    let newWorkout = {
      name
    };

    workouts.addWorkout({
      ...newWorkout,

      id: Math.random()
    });
    dispatch('addworkout')
  };
</script>

<Modal title="Add New Workout">
  <TextInput
    id="name"
    label="Name"
    value={name}
    valid={nameValid}
    validityMessage="Please enter a name"
    on:input={e => (name = e.target.value)} />
    <div slot="footer">
    <Button on:click={cancel}>Cancel</Button>
    <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
  </div>
</Modal>
