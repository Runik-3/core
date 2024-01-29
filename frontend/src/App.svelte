<script lang="ts">
  import DevicePanel from "./components/DevicePanel.svelte";
  import Generate from "./components/Generate.svelte";
  import Library from "./components/Library.svelte";
</script>

<main>
  <div id="gen">
    <Generate />
  </div>
  <div id="lib">
    <Library />
  </div>
  <div id="dev">
    <DevicePanel />
  </div>
</main>

<style>
  main {
    height: 100%;
    display: grid;
    grid-template-rows: 180px auto;
    grid-template-columns: auto minmax(200px, 300px);
    grid-template-areas:
      "gen gen"
      "lib dev";
  }
  #gen {
    grid-area: gen;
  }

  #lib {
    grid-area: lib;
  }

  #dev {
    grid-area: dev;
  }
</style>
