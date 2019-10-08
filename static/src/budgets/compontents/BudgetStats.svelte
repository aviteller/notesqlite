<script>
  import { afterUpdate } from "svelte";
  export let month;

  let totalIn = null;
  let totalOut = null;

  const getBudgetStats = month => {
    fetch(`http://localhost:9000/api/budgets/stats/${month}`)
      .then(res => {
        if (!res.ok) {
          throw new Error("Issue fetching meetups");
        }
        return res.json();
      })
      .then(data => {
        totalIn = data["in"];
        totalOut = data["out"];
      })
      .catch(err => {
        console.log(err);
      });
  };


  afterUpdate(() => {
    getBudgetStats(month);
  });
</script>

<style>

</style>

<div>totalin: {totalIn} totalin: {totalOut}</div>
