<script lang="ts">
  import { nav } from "./stores/nav";
  import DevicePanel from "./components/DevicePanel.svelte";
  import Generate from "./components/Generate.svelte";
  import Library from "./components/Library.svelte";
  import Navbar from "./components/Navbar.svelte";
  import NotificationProvider from "./components/NotificationProvider.svelte";
</script>

<main>
  <div id="nav">
    <Navbar />
  </div>
  <div id="view">
    <Library hide={$nav !== "lib"} />
    <Generate hide={$nav !== "gen"} />
  </div>
  <div id="dev">
    <DevicePanel />
  </div>
  <NotificationProvider />
</main>

<style>
  main {
    height: 100%;
    display: grid;
    grid-template-rows: 48px auto;
    grid-template-columns: auto minmax(300px, 350px);
    grid-template-areas:
      "nav nav"
      "view dev";
  }
  #nav {
    grid-area: nav;
  }

  #view {
    grid-area: view;
  }

  #dev {
    grid-area: dev;
  }
</style>
