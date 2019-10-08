<script>
 import { createEventDispatcher } from "svelte";
  import Select from "../../UI/Select.svelte";
  import Filter from "./BudgetFilter.svelte";
  import BudgetStats from "./BudgetStats.svelte";
  export let selectedMonth = null;
  const dispatch = createEventDispatcher();

  if (selectedMonth === null) {
    let d = new Date();
    selectedMonth = d.getMonth() + 1;
  }

  const months = [
    { label: "Jan", value: 1 },
    { label: "Feb", value: 2 },
    { label: "Mar", value: 3 },
    { label: "Apr", value: 4 },
    { label: "May", value: 5 },
    { label: "Jun", value: 6 },
    { label: "Jul", value: 7 },
    { label: "Aug", value: 8 },
    { label: "Sep", value: 9 },
    { label: "Oct", value: 10 },
    { label: "Nov", value: 11 },
    { label: "Dec", value: 12 }
  ];


  const handleSelect = (e) => {
    dispatch('changemonth', e.detail)
    selectedMonth = e.detail
  }
</script>

<style>
  .controls {
    display: flex;
    justify-content: space-between
  }

  .month {
    max-width: 30%;
  }
</style>

<div class="controls">

  <div class="month">
    <Select options={months} selected={parseInt(selectedMonth)} on:selectchange={handleSelect}/>
  </div>
  <div class="stats">
  <BudgetStats month={selectedMonth}/>
  </div>
  <div class="table-filter">
    <Filter on:select/>
  </div>
</div>
