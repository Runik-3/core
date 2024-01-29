<script lang="ts">
  import { nav, type Nav } from "../stores/nav";

  const navigate = (location: Nav) => {
    nav.set(location);
  };
</script>

<nav>
  <button on:click={() => navigate("lib")}>Lib</button>
  <button on:click={() => navigate("gen")}>Gen</button>
</nav>
