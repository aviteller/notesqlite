<script>
  import { createEventDispatcher } from "svelte";
  import Select from "../../UI/Select.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import { isEmpty } from "../../helpers/validation";
  import budgets from "../budgets-store.js";

  export let id = null;

  const formatDate = date => {
    let day = date.getDate();
    let month = date.getMonth() + 1;
    let year = date.getFullYear();
    if (day.toString().length === 1) {
      day = `0${day}`;
    }
    if (month.toString().length === 1) {
      month = `0${month}`;
    }
    return `${year}-${month}-${day}`;
  };

  let today = new Date();
  let name = "";
  let category = 2;
  let price = "";
  let date = formatDate(today);
  let type = 1;

  const types = [{ label: "IN", value: 0 }, { label: "OUT", value: 1 }];
  const categories = [
    { label: "wage", value: 1 },
    { label: "grocery", value: 2 },
    { label: "household", value: 3 }
  ];

  $: nameValid = !isEmpty(name);
  $: formIsValid = nameValid;

  const handleSelect = (e, field) => {
    // const som ={
    //   "category":category = e.detail,
    //   "type":type = e.detail,
    // }[field]
    switch (field) {
      case "category":
        category = e.detail;
        break;
      case "type":
        type = e.detail;
        break;

      default:
        break;
    }
  };

  const dispatch = createEventDispatcher();

  const submitForm = () => {
    const newBudget = {
      name,
      category: categories.find(({ value }) => value === category).label,
      price:parseFloat(price),
      date,
      type
    };

    fetch("http://localhost:9000/api/budgets", {
          method: "POST",
          body: JSON.stringify(newBudget),
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
             budgets.addBudget({...newBudget,  id: data.budget.id})
      
          })
          .catch(err => console.log(err));

    //budgets.addBudget({...newBudget, id:Math.random()})
  };
</script>

<style>
  form {
    display: flex;
    justify-content: space-between;
  }
</style>

<form action="">
  <TextInput
    id="name"
    label="Name"
    value={name}
    valid={nameValid}
    validityMessage="Please enter valid name"
    on:input={e => (name = e.target.value)} />
  <TextInput
    id="price"
    label="Price"
    value={price}
     float={true}
    type="number"
    on:input={e => (price = e.target.value)} />
  <TextInput
    id="date"
    label="Date"
    value={date}
    type="date"
    on:input={e => (date = e.target.value)} />
  <Select
    options={types}
    selected={type}
   
    on:selectchange={e => handleSelect(e, 'type')}
    label="Type" />
  <Select
    options={categories}
    selected={category}
    on:selectchange={e => handleSelect(e, 'category')}
    label="Category" />
  <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
</form>
