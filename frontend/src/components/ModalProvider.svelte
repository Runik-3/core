<script lang="ts">
  import { modalStore } from "../stores/modal";
  import Basic from "./modals/basic.svelte";
  import Dict from "./modals/dict.svelte";

  const { modalType } = $modalStore;
</script>

<div id="overlay"></div>
{#if modalType === "dict"}
  <Dict />
{:else}
  <Basic />
{/if}

<style>
  #overlay {
    position: absolute;
    top: 0;
    left: 0;
    background-color: black;
    opacity: 50%;
    height: 100vh;
    width: 100vw;
  }
</style>
