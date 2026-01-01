<script lang="ts">
  import Basic from "./basic.svelte";
  import { device } from "../../stores/device";
  import DictListItem from "../DictListItem.svelte";
  import {
    DeleteDeviceDictFile,
  } from "../../../wailsjs/go/main/App";
</script>

<Basic>
  {#if $device?.dicts?.length}
    <ul id="device-dict-list">
      {#each $device.dicts as dict}
        <DictListItem
          isDeviceList
          {dict}
          toggleChecked={() => {}}
          selectDict={() => {}}
          deleteDict={DeleteDeviceDictFile}
        />
      {/each}
    </ul>
  {:else}
    No dictionaries found on this device.
  {/if}
</Basic>

<style>
  #device-dict-list {
    box-sizing: border-box;
    background-color: var(--bg);
    margin-top: 16px;
    border: 1px #a5a5a5 solid;
    border-radius: 8px;
    padding: 8px 16px;
    min-height: 12rem;
    max-height: 18rem;
    margin-bottom: 1rem;
    overflow-y: auto;
  }
</style>
