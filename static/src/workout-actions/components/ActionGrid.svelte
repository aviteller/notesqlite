<script>
  import ActionItem from "./ActionItem.svelte";
  import EditAction from "./EditAction.svelte";
  import Button from "../../UI/Button.svelte";
  import { onDestroy, createEventDispatcher, onMount } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";
  import SortableList from "../../UI/SortableList.svelte";

  export let actions;

  const dispatch = createEventDispatcher();

  const sortList = ev => {
 
    actions = ev.detail;
  };

  const saveSort = (e) => {
    fetch(`http://localhost:9000/api/actions/swap/${e.detail[0]}/${e.detail[1]}`)
    .then(res => {
      if (!res.ok) {
        throw new Error("Issue fetching workous");
      }
      return res.json();
    })
    .then(data => {
      console.log(data)
    })
    .catch(err => {
      isLoading = false;
      console.log(err);
    });
  } 

</script>

<style>
  #meetups {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr;
    grid-gap: 1rem;
  }

  @media (min-width: 768px) {
    #meetups {
      grid-template-columns: repeat(4, 1fr);
    }
  }

</style>


<section id="meetups">
  {#if actions && actions.length > 0}
    <SortableList list={actions} key="id" on:sort={sortList} let:item on:savesort={saveSort}>
      <ActionItem
        id={item.id}
        name={item.name}
        action_length={item.action_length}
        action_type={item.action_type}
        equipment={item.equipment}
        workoutID={item.workoutID}
        on:edit />
    </SortableList>
    <!-- {#each $actions as action (action.id)}
          <div transition:scale animate:flip={{ duration: 300 }}>
            <ActionItem
              id={action.id}
              name={action.name}
              action_length={action.action_length}
              action_type={action.action_type}
              equipment={action.equipment}
              workoutID={action.workoutID}
              on:edit={starEdit} />
          </div>
        {/each} -->
  {:else}
    <h1>Add actions</h1>
  {/if}
</section>
