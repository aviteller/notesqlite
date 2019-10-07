<script>
  import BudgetTable from "./compontents/BudgetTable.svelte";
  import BudgetForm from "./compontents/BudgetForm.svelte";
  import budgets from "./budgets-store.js";

  let isLoading = true;

    fetch(`http://localhost:9000/api/budgets`)
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
          const loadedBudgets = [];
          for (const key in data.data) {
            loadedBudgets.push({
              ...data.data[key]
            });
          }
          budgets.setBudgets(loadedBudgets);
          isLoading = false;
        }
      })
      .catch(err => {
        error = err;
        isLoading = false;
      });

</script>

<BudgetForm />
{#if !isLoading}
<BudgetTable budgets={$budgets} />
{/if}
