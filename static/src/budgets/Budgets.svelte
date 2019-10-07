<script>
  import BudgetTable from "./compontents/BudgetTable.svelte";

  import budgets from "./budgets-store.js";


</script>


<BudgetTable budgets={$budgets} />
