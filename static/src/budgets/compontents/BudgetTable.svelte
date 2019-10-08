<script>
  import BudgetItem from "./BudgetItem.svelte";
  import BudgetControls from "./BudgetControls.svelte";
  import Table from "../../UI/Table.svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  export let budgets;

  let filter = false;

  const groupBy = (array, fn) =>
    array.reduce((result, item) => {
      const key = fn(item);
      if (!result[key]) result[key] = [];
      result[key].push(item);
      return result;
    }, []);

  $: filteredBudgets = filter
    ? filter === 1
      ? budgets.filter(b => b.type === 0)
      : filter === 2
      ? budgets.filter(b => b.type === 1)
      : filter === 3
      ? groupBy(budgets, b => (b.type === 0 ? 0 : 1))
      : budgets
    : budgets;

  const setFilter = e => {
    filter = e.detail;
  };

  const tableHeaders = ["ID", "Name", "Category", "Price", "Type", "Date", ""];
</script>

<style>
  #budget-table {
    width: 100%;
  }
</style>

<BudgetControls on:select={setFilter} on:changemonth/>
<section id="budget-table">

  {#if filteredBudgets}
    {#if filter !== 3}
      <Table headers={tableHeaders}>
        {#each filteredBudgets as budget (budget.id)}
          <BudgetItem {budget} on:edit/>
        {/each}
      </Table>
    {:else}
      {#each filteredBudgets as filteredBudget, i}
        {#if filteredBudget}
          <Table headers={tableHeaders} title={i == 0 ? 'IN' : 'OUT'}>
            {#each filteredBudget as budget (budget.id)}
              <BudgetItem {budget} on:edit/>
            {/each}
          </Table>
          <hr />
        {/if}
      {/each}
    {/if}
  {:else}
    <h1>Add new budgets</h1>
  {/if}
</section>
