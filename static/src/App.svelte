<script>
  import Header from "./UI/Header.svelte";
  import routes from "./routes";
  import Router from "svelte-spa-router";

  let title = "notes";
</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>

<Header>{title}</Header>
<main>
  <Router {routes} />
</main>
