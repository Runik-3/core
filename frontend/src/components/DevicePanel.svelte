<script lang="ts">
  import {
    SelectDevice,
    DeleteDeviceDictFile,
  } from "../../wailsjs/go/main/App";
  import Button from "./Button.svelte";
  import Device from "./icons/Device.svelte";
  import { device } from "../stores/device";
  import DictListItem from "./DictListItem.svelte";

  const openDir = async () => {
    const devicePath = await SelectDevice();
    await device.selectDevice(devicePath);
  };
</script>

<div id="device-panel">
  <h2>Device</h2>
  <div class="content">
    {#if $device.path}
      <Device size="128px" color="#5D5D5D" />
      <div id="device-info">
        <p>
          <strong>Name:</strong>
          {$device.name}
        </p>
        <!-- List device dictionaries -->
        <strong>Dictionaries:</strong>
        {#if $device?.dicts?.length}
          {#each $device.dicts as dict}
            <DictListItem
              compact
              {dict}
              select={() => {}}
              deleteDict={DeleteDeviceDictFile}
            />
          {/each}
        {:else}
          No dictionaries found on this device.
        {/if}
      </div>
    {:else}
      <Device size="128px" hasPlug color="#5D5D5D" />
      <p>Please plug-in and select your reader device.</p>
    {/if}
  </div>
  <Button onClick={openDir} maxWidth small>Select Device</Button>
</div>

<style>
  #device-panel {
    padding: 24px;
    height: calc(100% - 48px); /* minus top nav */
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .content {
    text-align: center;
  }
  #device-info {
    text-align: left;
  }
  #device-info > p {
    margin: 12px 0;
  }
  p {
    margin: 24px 0;
  }
</style>
