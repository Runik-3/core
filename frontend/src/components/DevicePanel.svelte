<script lang="ts">
  import type { DictFile } from "../types/dict";
  import {
    SelectDevice,
    GetDeviceDictionaries,
  } from "../../wailsjs/go/main/App";
  import Button from "./Button.svelte";
  import Device from "./icons/Device.svelte";
  import { Severity, notifications } from "../stores/notification";
  import type { Response } from "../types/response";
  import DictListItem from "./DictListItem.svelte";

  let devicePath = "";
  let dicts: DictFile[] = [];

  $: if (devicePath) getDeviceDicts();

  const openDir = async () => {
    devicePath = await SelectDevice();
  };

  const getDeviceNameFromPath = (path: string) => {
    const parts = path.split("/");
    return parts[parts.length - 1];
  };

  const getDeviceDicts = async () => {
    const res: Response<DictFile[]> = await GetDeviceDictionaries();
    if (res.Error) {
      notifications.addNotificaton({
        message: `Error fetching device dictionaries.\n${res.Error}`,
        severity: Severity.error,
      });
    }
    dicts = res.Data.filter((file) => file.Extension === "zip");
  };
  $: console.log(dicts);
</script>

<div id="device-panel">
  <h2>Device</h2>
  <div class="content">
    {#if devicePath}
      <Device size="128px" color="#5D5D5D" />
      <div id="device-info">
        <p>
          <strong>Name:</strong>
          {getDeviceNameFromPath(devicePath)}
        </p>
        <!-- List device dictionaries -->
        <strong>Dictionaries:</strong>
        {#if dicts.length}
          {#each dicts as dict}
            <DictListItem compact {dict} select={() => {}} />
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
