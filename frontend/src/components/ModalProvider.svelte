<script lang="ts">
  import { modalStore } from "../stores/modal";
  import Basic from "./modals/basic.svelte";
  import ConvertSelect from "./modals/convertSelect.svelte";
  import RenameDict from "./modals/renameDict.svelte";
  import Device from "./modals/device.svelte";

  const { modalType } = $modalStore;
</script>

<div id="overlay"></div>
{#if modalType === "convertSelect"}
  <ConvertSelect />
{:else if modalType === "renameDict"}
  <RenameDict />
{:else if modalType === "device"}
  <Device/>
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
