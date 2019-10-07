<script>
  import BudgetItem from "./BudgetItem.svelte";
  import BudgetControls from "./BudgetControls.svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  export let budgets;

  let filter = false;

  $: [filteredBudgets, filteredBudgets2] = filter
    ? filter === 1
      ? budgets.reduce(
          (output, b) => {
            if (b.type === "in") output[0].push(b);
            //if (b.type === "out") output[1].push(b);
            return output;
          },
          [[], []]
        )
      : filter === 2
      ? budgets.reduce(
          (output, b) => {
            if (b.type === "out") output[0].push(b);
            //if (b.type === "in") output[1].push(b);

            return output;
          },
          [[], []]
        )
      : filter === 3
      ? budgets.reduce(
          (output, b) => {
            if (b.type === "out") output[0].push(b);
            if (b.type === "in") output[1].push(b);

            return output;
          },
          [[], []]
        )
      : budgets.reduce(
          (output, b) => {
            output[0].push(b);

            return output;
          },
          [[], []]
        )
    : budgets.reduce(
        (output, b) => {
          output[0].push(b);

          return output;
        },
        [[], []]
      );

  const setFilter = e => {
    filter = e.detail;
    console.log(filteredBudgets);
  };
</script>

<style>
  #budget-table {
    width: 100%;
  }

  table {
    width: 100%;
  }
</style>

<BudgetControls on:select={setFilter} />
<section id="budget-table">

  {#if filteredBudgets.length > 0}
    {#if filter !== 3}
      <table>
        <thead>
          <tr>
            <td>ID</td>
            <td>Name</td>
            <td>Category</td>
            <td>Price</td>
            <td>Type</td>
            <td>Date</td>
            <td />
          </tr>
        </thead>
        <tbody>
          {#each filteredBudgets as budget (budget.id)}
            <BudgetItem {budget} />
          {/each}
        </tbody>
      </table>
    {:else}
      <table>
        <thead>
          <tr>
            <td>ID</td>
            <td>Name</td>
            <td>Category</td>
            <td>Price</td>
            <td>Type</td>
            <td>Date</td>
            <td />
          </tr>
        </thead>
        <tbody>
          {#each filteredBudgets as budget (budget.id)}
            <BudgetItem {budget} />
          {/each}
        </tbody>
      </table>
      <table>
        <thead>
          <tr>
            <td>ID</td>
            <td>Name</td>
            <td>Category</td>
            <td>Price</td>
            <td>Type</td>
            <td>Date</td>
            <td />
          </tr>
        </thead>
        <tbody>
          {#each filteredBudgets2 as budget (budget.id)}
            <BudgetItem {budget} />
          {/each}
        </tbody>
      </table>
    {/if}
  {:else}
    <h1>Add new budgets</h1>
  {/if}
</section>
