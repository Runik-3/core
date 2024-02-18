<script lang="ts">
  import { SelectDevice } from "../../wailsjs/go/main/App";
  import Button from "./Button.svelte";
  import Device from "./icons/Device.svelte";
  let devicePath = "";

  const openDir = async () => {
    devicePath = await SelectDevice();
  };

  const getDeviceFromPath = (path: string) => {
    const parts = path.split("/");
    return parts[parts.length - 1];
  };
</script>

<div id="device-panel">
  <h2>Device</h2>
  <div class="content">
    {#if devicePath}
      <Device size="128px" color="#5D5D5D" />
      <div id="device-info">
        <p>
          <strong>Name:</strong>
          {getDeviceFromPath(devicePath)}
        </p>
        <!-- List device dictionaries -->
      </div>
    {:else}
      <p>Please select a reader device.</p>
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
  h2 {
    margin-bottom: 24px;
  }
  p {
    margin: 24px 0;
  }
</style>
