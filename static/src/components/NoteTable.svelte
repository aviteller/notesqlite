<script>
  import NoteItem from './NoteItem.svelte'
  export let notes;

</script>


<style>
.notes-table {width:100%}
</style>

<table class="notes-table">
  <thead>
    <tr>
      <th>ID</th>
      <th>Title</th>
      <th style="width:50%">Message</th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    {#each notes as note (note.id)}
      <tr>
        <NoteItem ID={note.id} title={note.title} message={note.message} on:edit/>
      </tr>
    {/each}
  </tbody>
</table>