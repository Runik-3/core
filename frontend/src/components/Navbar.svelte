<script lang="ts">
  import { nav, type Nav } from "../stores/nav";
  import Anvil from "./icons/Anvil.svelte";
  import Book from "./icons/Book.svelte";
  import Runik from "./icons/Runik.svelte";
  import Cog from "./icons/Cog.svelte";

  const navigate = (location: Nav) => {
    nav.set(location);
  };
</script>

<nav>
  <button on:click={() => {}}><Runik color="white" size="24px" /></button>
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
    align-items: center;
  }
  button {
    width: 72px;
    height: 100%;
    background-color: transparent;
    border: none;
    cursor: pointer;
  }
  button:first-of-type {
    background-color: #202020;
  }
  .active {
    background-color: white;
  }
</style>
