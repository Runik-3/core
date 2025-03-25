<script lang="ts">
  import { nav, type Nav } from "../stores/nav";
  import Anvil from "./icons/Anvil.svelte";
  import Book from "./icons/Book.svelte";
  import Cog from "./icons/Cog.svelte";

  const navigate = (location: Nav) => {
    nav.set(location);
  };
</script>

<nav>
  <button
    class={`${$nav === "lib" && "active"}`}
    on:click={() => navigate("lib")}
    ><Book size="20px" color={`${$nav === "lib" ? "black" : "#5D5D5D"}`} /><span
      >Library</span
    ></button
  >
  <button
    class={`${$nav === "gen" && "active"}`}
    on:click={() => navigate("gen")}
    ><Anvil
      size="20px"
      color={`${$nav === "gen" ? "black" : "#5D5D5D"}`}
    /><span>Forge</span></button
  >
  <button
    class={`${$nav === "conf" && "active"}`}
    on:click={() => navigate("conf")}
    ><Cog size="20px" color={`${$nav === "conf" ? "black" : "#5D5D5D"}`} /><span
      >Settings</span
    ></button
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
    padding: 0 12px;
    margin-left: 0.5rem;
    border-radius: 4px;
    height: 80%;
    background-color: transparent;
    border: none;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #5d5d5d;
  }
  button:hover {
    background-color: #b5b5b5;
  }
  .active {
    background-color: white;
    color: black;
  }
  .active:hover {
    background-color: white;
  }
  button span {
    font-size: 0.9rem;
    margin-left: 8px;
  }
</style>
