<script>
  import { createEventDispatcher } from "svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import Modal from "../../UI/Modal.svelte";
  import { isEmpty, isValidEmail } from "../../helpers/validation";
  import workouts from "../workouts-store.js";
  let name = "";
  let duration = "";
  let workout_type = "";
  const dispatch = createEventDispatcher();
  const cancel = () => dispatch("cancel");

  $: nameValid = !isEmpty(name);
  $: formIsValid = nameValid;

  const submitForm = () => {
    let newWorkout = {
      name,
      duration: parseInt(duration),
      workout_type
    };

    console.log(newWorkout);

    fetch("http://localhost:9000/api/workouts", {
      method: "POST",
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
        workouts.addWorkout({ ...newWorkout, id: data.workout.id });
        dispatch("addworkout");
      })
      .catch(err => console.log(err));

    // workouts.addWorkout({
    //   ...newWorkout,

    //   id: Math.random()
    // });
    // dispatch('addworkout')
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

  <TextInput
    id="duration"
    label="Duration"
    value={duration}
    type="number"
    on:input={e => (duration = e.target.value)} />

  <TextInput
    id="workout_type"
    label="Workout Type"
    value={workout_type}
    on:input={e => (workout_type = e.target.value)} />
  <div slot="footer">
    <Button on:click={cancel}>Cancel</Button>
    <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
  </div>
</Modal>
