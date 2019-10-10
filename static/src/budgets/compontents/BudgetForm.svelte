<script>
  import { createEventDispatcher, onMount } from "svelte";
  import Select from "../../UI/Select.svelte";
  import TextInput from "../../UI/TextInput.svelte";
  import Button from "../../UI/Button.svelte";
  import { isEmpty } from "../../helpers/validation";
  import budgets from "../budgets-store.js";
  import categories from "../budget-categories-store.js";

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
  let category = 0;
  let categoryLabel = "";
  let price = "";
  let date = formatDate(today);
  let type = 1;

  const types = [{ label: "IN", value: 0 }, { label: "OUT", value: 1 }];

  const getCategories = (type = 1) => {
    fetch(`http://localhost:9000/api/budgetcategories/${type}`)
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
          const loadedCategories = [];
          for (const key in data.data) {
            loadedCategories.push({
              ...data.data[key]
            });
          }
          categories.setBudgetCategories(loadedCategories);
        }
      })
      .catch(err => {
        console.log(err);
      });
  };

  onMount(() => {
    getCategories(1);
  });

  if (id) {
    const unsubscribe = budgets.subscribe(items => {
      const selectedBudget = items.find(i => i.id === id);
      //console.log(selectedBudget)
      if (selectedBudget) {
        name = selectedBudget.name;
        (category = categories.find(
          ({ label }) => label === selectedBudget.category
        ).value),
          (price = selectedBudget.price);
        date = selectedBudget.date;
        type = selectedBudget.type;
      }
    });

    unsubscribe();
  }

  $: nameValid = !isEmpty(name);
  $: formIsValid = nameValid;

  const handleSelect = (e, field) => {
    
    switch (field) {
      case "category":
        category = e.detail.selectedValue;
        categoryLabel = e.detail.selectedLabel;
        break;
      case "type":
        type = e.detail.selectedValue;
        getCategories(e.detail.selectedValue)
    }
  };

  const dispatch = createEventDispatcher();

  const submitForm = () => {
    const newBudget = {
      name,
      category:category.toString(),
      price: parseFloat(price),
      date,
      type
    };

    
    if (id) {
      fetch(`http://localhost:9000/api/budgets/${id}`, {
        method: "PUT",
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
          budgets.updateBudget(id, newBudget);
          cancel();
        })
        .catch(err => console.log(err));
    } else {
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

          cancel();
         
          budgets.addBudget({ ...newBudget, id: data.budget.id,  category:categoryLabel});
        })
        .catch(err => console.log(err));
    }
    //budgets.addBudget({...newBudget, id:Math.random()})
  };

  const cancel = () => {
    dispatch("cancel");
  };
</script>

<style>
  form {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
</style>

<form action="" class="form">
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
    selectID="type"
    on:selectchange={e => handleSelect(e, 'type')}
    label="Type" />
  <Select
    options={$categories}
    selected={category}
    selectID="categories"
    disabled={false}
    on:selectchange={e => handleSelect(e, 'category')}
    label="Category" />

  <Button on:click={submitForm} disabled={!formIsValid}>Save</Button>
  {#if id}
    <Button color="danger" on:click={cancel}>Cancel</Button>
  {/if}
</form>
