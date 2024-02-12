<script lang="ts">
  import { SelectDevice } from "../../wailsjs/go/main/App";
  import Button from "./Button.svelte";
  import Device from "./icons/Device.svelte";
  let devicePath = "";

  const openDir = async () => {
    devicePath = await SelectDevice();
  };
</script>

<div>
  <h3>Devices</h3>
  <div class="content">
    <Device size="128px" color="#5D5D5D" />
    <p>
      {devicePath
        ? `Device Selected: ${devicePath}`
        : "Please choose your e-reader device."}
    </p>
    <Button onClick={openDir} small>Select Reader Device</Button>
  </div>
</div>

<style>
  div {
    padding: 24px;
  }
  .content {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  h3 {
    margin-bottom: 24px;
  }
  p {
    margin: 24px 0;
  }
</style>
