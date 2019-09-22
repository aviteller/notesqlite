<!-- <script>
  import { createEventDispatcher } from "svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import Modal from "../../UI/Modal.svelte";
  import { isEmpty, isValidEmail } from "../../helpers/validation";
   import workouts from "../workouts-store.js";

  export let id = null;

  let name = "";
  let subtitle = "";
  let contactEmail = "";
  let description = "";
  let imageUrl = "";
  let address = "";

  if (id) {
    const unsubscribe = meetups.subscribe(items => {
      const selectedMeetup = items.find(i => i.id === id);
      title = selectedMeetup.title;
      subtitle = selectedMeetup.subtitle;
      contactEmail = selectedMeetup.contactEmail;
      description = selectedMeetup.description;
      imageUrl = selectedMeetup.imageUrl;
      address = selectedMeetup.address;
    });

    unsubscribe();
  }

  $: titleValid = !isEmpty(title);
  $: subtitleValid = !isEmpty(subtitle);
  $: contactEmailValid = isValidEmail(contactEmail);
  $: descriptionValid = !isEmpty(description);
  $: imageUrlValid = !isEmpty(imageUrl);
  $: addressValid = !isEmpty(address);
  $: formIsValid =
    titleValid &&
    subtitleValid &&
    contactEmailValid &&
    descriptionValid &&
    imageUrlValid &&
    addressValid;

  const dispatch = createEventDispatcher();

  const clearForm = () => {
    title = "";
    subtitle = "";
    contactEmail = "";
    description = "";
    imageUrl = "";
    address = "";
  };

  //   const submitForm = e => {

  //   editMode = null;
  // };

  const submitForm = () => {
    const newMeetup = {
      title,
      subtitle,
      contactEmail,
      description,
      imageUrl,
      address
    };
    if (id) {
      fetch(`http://localhost:9000/api/meetups/${id}`, {
        method: "PUT",
        body: JSON.stringify(newMeetup),
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
          meetups.updateMeetup(id, newMeetup);
          dispatch("addmeetup");
          clearForm();
        })
        .catch(err => console.log(err));
      meetups.updateMeetup(id, newMeetup);
    } else {
      fetch("http://localhost:9000/api/meetups", {
        method: "POST",
        body: JSON.stringify(newMeetup),
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
          meetups.addMeetup({
            ...newMeetup,
            isLiked: false,
            id: data.meetup.ID
          });
          dispatch("addmeetup");
          clearForm();
        })
        .catch(err => console.log(err));
    }
  };
  const deleteMeetup = () => {
    fetch(`http://localhost:9000/api/meetups/${id}`, {
      method: "DELETE"
    })
      .then(res => {
        meetups.removeMeetup(id);
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

<Modal title="Edit Meetup" on:cancel>
  <form>
    <TextInput
      id="title"
      label="Title"
      value={title}
      valid={titleValid}
      validityMessage="Please enter valid title"
      on:input={e => (title = e.target.value)} />
    <TextInput
      id="subtitle"
      label="Subtitle"
      value={subtitle}
      valid={subtitleValid}
      validityMessage="Please enter valid subtitle"
      on:input={e => (subtitle = e.target.value)} />
    <TextInput
      id="address"
      label="Address"
      value={address}
      valid={addressValid}
      validityMessage="Please enter valid address"
      on:input={e => (address = e.target.value)} />

    <TextInput
      id="imageUrl"
      label="Image URL"
      value={imageUrl}
      valid={imageUrlValid}
      validityMessage="Please enter valid Image url"
      on:input={e => (imageUrl = e.target.value)} />
    <TextInput
      id="contactEmail"
      label="Email"
      value={contactEmail}
      valid={contactEmailValid}
      validityMessage="Please enter valid email"
      type="email"
      on:input={e => (contactEmail = e.target.value)} />
    <TextInput
      id="description"
      label="Description"
      value={description}
      valid={descriptionValid}
      validityMessage="Please enter valid description"
      type="textarea"
      rows="3"
      bind:value={description} />
  </form>
  <div slot="footer">
    <Button on:click={cancel}>Cancel</Button>
    <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
    {#if id}
      <Button on:click={deleteMeetup}>Delete</Button>
    {/if}
  </div>
</Modal> -->
