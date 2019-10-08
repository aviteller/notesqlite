<script>
  import { createEventDispatcher } from "svelte";

  export let options;
  export let label = null;
  export let placeholder = "--Select--";
  export let selected = null;

  const dispatch = createEventDispatcher();

  const handleChange = () => {
    dispatch("selectchange", selected);
  };
</script>

<style>
  /* label {
    margin-bottom: 10px;
  } */

  label {
    display: block;
    margin-bottom: 0.5rem;
    width: 100%;
  }
  select {
    display: block;
    width: 100%;
  }

  .form-control {
    padding: 0.5rem 0;
    width: 100%;
    margin: 0.25rem 1rem;
  }
</style>

<div class="form-control">

  {#if label}
    <label>{label}</label>
  {/if}
  <select bind:value={selected} on:change={handleChange}>
    <option value="" disabled>{placeholder}</option>
    {#each options as option}
      <option value={option.value}>{option.label}</option>
    {/each}
  </select>
</div>
