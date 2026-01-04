<script lang="ts">
  import Basic from "./basic.svelte";
  import { device } from "../../stores/device";
  import DictListItem from "../DictListItem.svelte";
  import { DeleteDeviceDictFile } from "../../../wailsjs/go/main/App";
</script>

<Basic>
  {#if $device?.dicts?.length}
    <h3>Dictionaries on this device</h3>
    <ul id="device-dict-list">
      {#each $device.dicts as dict}
        <DictListItem
          isDeviceList
          {dict}
          deleteDict={DeleteDeviceDictFile}
        />
      {/each}
    </ul>
  {:else}
    No dictionaries found on this device.
  {/if}
</Basic>

<style>
  h3 {
    margin-top: 1rem;
  }
  #device-dict-list {
    box-sizing: border-box;
    margin-top: 16px;
    border: 1px var(--bg-secondary) solid;
    border-radius: 8px;
    padding: 8px 16px;
    min-height: 12rem;
    max-height: 18rem;
    margin-bottom: 1rem;
    overflow-y: auto;
  }
</style>
