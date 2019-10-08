<script>
  import BudgetTable from "./compontents/BudgetTable.svelte";
  import BudgetForm from "./compontents/BudgetForm.svelte";
  import budgets from "./budgets-store.js";

  let isLoading = true;
  let editMode = false;
  let editId = null;
  let month = new Date().getMonth() + 1;

  const changeMonth = e => {
    getBudgets(e.detail);
  };

  const getBudgets = month => {
    fetch(`http://localhost:9000/api/budgets/${month}`)
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
  };

  getBudgets(month);

  const startEdit = e => {
    cancelEdit();
    setTimeout(() => {
      editMode = true;
      editId = e.detail;
    }, 500);
  };

  const cancelEdit = () => {
    editMode = false;
    editId = null;
  };
</script>

{#if editMode}
  <BudgetForm id={editId} on:cancel={cancelEdit} />
{:else}
  <BudgetForm />
{/if}
{#if !isLoading}
  <BudgetTable
    budgets={$budgets}
    on:edit={startEdit}
    on:changemonth={changeMonth} />
{/if}
