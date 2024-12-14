<script lang="ts">
  import { nav, type Nav } from "../stores/nav";
  import Anvil from "./icons/Anvil.svelte";
  import Book from "./icons/Book.svelte";

  const navigate = (location: Nav) => {
    nav.set(location);
  };
</script>

<nav>
  <button
    class={`${$nav === "lib" && "active"}`}
    on:click={() => navigate("lib")}
    ><Book
      size="24px"
      color={`${$nav === "lib" ? "black" : "#5D5D5D"}`}
    /></button
  >
  <button
    class={`${$nav === "gen" && "active"}`}
    on:click={() => navigate("gen")}
    ><Anvil
      size="32px"
      color={`${$nav === "gen" ? "black" : "#5D5D5D"}`}
    /></button
  >
  <!-- Disabled until we need global config -->
  <!-- <button
    class={`${$nav === "conf" && "active"}`}
    on:click={() => navigate("conf")}
    ><Cog
      size="24px"
      color={`${$nav === "conf" ? "black" : "#5D5D5D"}`}
    /></button
  > -->
</nav>

<style>
  nav {
    height: 100%;
    display: flex;
    justify-content: start;
    align-items: end;
  }
  button {
    width: 72px;
    border-radius: 4px 4px 0 0;
    height: 80%;
    background-color: transparent;
    border: none;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  button:hover {
    background-color: #b5b5b5;
  }
  .active {
    background-color: white;
  }
  .active:hover {
    background-color: white;
  }
</style>
