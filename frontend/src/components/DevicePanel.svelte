<script lang="ts">
  import { SelectDevice } from "../../wailsjs/go/main/App";
  let devicePath = "";

  const openDir = async () => {
    devicePath = await SelectDevice();
  };
</script>

<div>
  <button on:click={openDir}>Select Reader Device</button> <br />
  <p>
    {devicePath
      ? `Device Selected: ${devicePath}`
      : "Please choose your e-reader device."}
  </p>
</div>
