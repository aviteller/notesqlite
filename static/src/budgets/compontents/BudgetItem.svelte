<script>
  import Button from "../../UI/Button.svelte";

  export let budget;

  const { id, name, category, price, type, date } = budget;
</script>

<tr>
  <td>{id}</td>
  <td>{name}</td>
  <td>{category}</td>
  <td>{price}</td>
  <td>{type?'OUT':'IN'}</td>
  <td>{date}</td>
  <td>
    <Button>Edit</Button>
    <Button color="danger">X</Button>
  </td>
</tr>
