<script lang="ts">
  import {
    SelectDevice,
    DeleteDeviceDictFile,
  } from "../../wailsjs/go/main/App";
  import Button from "./Button.svelte";
  import Device from "./icons/Device.svelte";
  import { device } from "../stores/device";
  import DictListItem from "./DictListItem.svelte";
  import type { Device as DeviceType } from "../types/device";
  import type { Response } from "../types/response";
  import { notifications, Severity } from "../stores/notification";

  const openDir = async () => {
    const deviceRes: Response<DeviceType> = await SelectDevice();
    if (deviceRes.Error) {
      notifications.addNotification({
        message: deviceRes.Error,
        severity: Severity.warn,
        timeout: 5000,
      });
    }
    await device.selectDevice(deviceRes.Data.Path);
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
          <ul id="device-dict-list">
            {#each $device.dicts as dict}
              <DictListItem
                compact
                {dict}
                toggleSelect={() => {}}
                deleteDict={DeleteDeviceDictFile}
              />
            {/each}
          </ul>
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
    overflow: auto;
    display: grid;
    grid-template-rows: 48px auto 32px;
  }
  .content {
    margin-top: 48px;
    text-align: center;
  }
  #device-info {
    text-align: left;
  }
  #device-info > p {
    margin: 12px 0;
  }
  #device-dict-list {
    background-color: white;
    margin-top: 16px;
    border: 1px #a5a5a5 solid;
    border-radius: 8px;
    padding: 8px 16px;
    height: calc(100vh - 524px); /* overflow needs a set number for height */
    overflow-y: auto;
  }
  p {
    margin: 24px 0;
  }
</style>
