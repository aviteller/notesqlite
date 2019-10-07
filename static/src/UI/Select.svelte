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
  label {
    margin-bottom: 10px;
  }
  select {
    width: 100%;
  }
</style>

{#if label}
  <label>{label}</label>
{/if}
<select bind:value={selected} on:change={handleChange}>
  <option value="" disabled>{placeholder}</option>
  {#each options as option}
    <option value={option.value}>{option.label}</option>
  {/each}
</select>
