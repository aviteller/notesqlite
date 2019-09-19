<script>
  import NavLink from "./Navlink.svelte";
  import Button from "./Button.svelte";
</script>

<style>
  header {
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    height: 4rem;
    background-color: white;
    border-bottom: 12px solid #199b14;
    display: flex;
    justify-content: space-evenly;
    align-content: center;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.26);
    margin-bottom: 20px;
  }

  h1 {
    color: black;
    margin-top: 5px;
    text-transform: uppercase;
  }

  .links {
    margin-top: 1rem;
  }
</style>

<header>
  <h1>
    <slot />
  </h1>
  <div class="links">
    <Button href="/workouts">Workouts</Button>
    <Button href="/">Notes</Button>

  </div>
</header>
