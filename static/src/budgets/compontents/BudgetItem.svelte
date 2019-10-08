<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../../UI/Button.svelte";
  import budgets from "../budgets-store.js";

  export let budget;

  const dispatch = createEventDispatcher();

  const { id, name, category, price, type } = budget;
  const date = budget.date.split("T")[0];

  const removeBudget = id => {
    if (window.confirm(`Are you sure you want to delete: ${name}`)) {
      fetch(`http://localhost:9000/api/budgets/${id}`, {
        method: "DELETE"
      })
        .then(res => {
          budgets.removeBudget(id);
        })
        .catch(err => console.log(err));
    }
  };
</script>

<tr>
  <td>{id}</td>
  <td>{name}</td>
  <td>{category}</td>
  <td>{price}</td>
  <td>{type ? 'OUT' : 'IN'}</td>
  <td>{date}</td>
  <td>
    <Button on:click={() => dispatch('edit', id)}>Edit</Button>
    <Button color="danger" on:click={() => removeBudget(id)}>X</Button>
  </td>
</tr>
