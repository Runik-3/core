<script lang="ts">
  import { nav, type Nav } from "../stores/nav";
  import Anvil from "./icons/Anvil.svelte";
  import Book from "./icons/Book.svelte";
  import Runik from "./icons/Runik.svelte";

  const navigate = (location: Nav) => {
    nav.set(location);
  };
</script>

<nav>
  <button on:click={() => navigate("lib")}><Runik size="24px" /></button>
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
</nav>

<style>
  nav {
    height: 100%;
    display: flex;
    justify-content: start;
    align-items: center;
  }
  button {
    width: 72px;
    height: 100%;
    background-color: transparent;
    border: none;
    cursor: pointer;
  }
  .active {
    background-color: white;
  }
</style>
