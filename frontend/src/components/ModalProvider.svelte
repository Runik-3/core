<script lang="ts">
  import { modalStore } from "../stores/modal";
  import Basic from "./modals/basic.svelte";
  import ConvertSelect from "./modals/convertSelect.svelte";
  import RenameDict from "./modals/renameDict.svelte";
  import Device from "./modals/device.svelte";
  import Details from "./modals/details.svelte";
</script>

<div id="overlay"></div>
{#if $modalStore.modalType === "convertSelect"}
  <ConvertSelect />
{:else if $modalStore.modalType === "renameDict"}
  <RenameDict />
{:else if $modalStore.modalType === "device"}
  <Device />
{:else if $modalStore.modalType === "details"}
  <Details />
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
