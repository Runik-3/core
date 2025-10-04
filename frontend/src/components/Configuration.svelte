<script lang="ts">
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import InfoPopover from "./InfoPopover.svelte";
  import { GetConfig, SelectFile, SetConfig } from "../../wailsjs/go/main/App";
  import type { Response } from "../types/response";
  import type { Config } from "../types/config";
  import { notifications, Severity } from "../stores/notification";
  import { onMount } from "svelte";
  Error;

  export let hide = false;
  let config: Config | undefined = undefined;

  onMount(async () => {
    const res: Response<Config> = await GetConfig();
    if (res.Error) {
      notifications.addNotification({
        message: res.Error,
        severity: Severity.error,
        timeout: 5000,
      });
    }
    config = res.Data;
  });

  const handleSelectFile = async () => {
    const res: Response<string> = await SelectFile({});
    if (res.Error) {
      notifications.addNotification({
        message: res.Error,
        severity: Severity.warn,
        timeout: 5000,
      });
      return;
    }
    // Return file path
    return res.Data;
  };

  const setKindlegenValue = async (value?: string) => {
    // TODO: loading state while config saves
    // This func has no return value
    config["kindlegenPath"] = value;
    const setRes: Response<string> = await SetConfig(config);
    if (setRes.Error) {
      notifications.addNotification({
        message: setRes.Error,
        severity: Severity.info,
        timeout: 5000,
      });
    }

    // Fetch updated new config
    const getRes: Response<Config> = await GetConfig();
    if (getRes.Error) {
      notifications.addNotification({
        message: getRes.Error,
        severity: Severity.error,
        timeout: 5000,
      });
    }

    return getRes.Data;
  };

  /** Accept value arg  */
  const handleSetKindlegenPathFromFileSelect = async (value?: string) => {
    const newPath = await handleSelectFile();
    if (newPath) {
      const newConfig = await setKindlegenValue(newPath);
      config = newConfig;
    }
  };

  const handleResetKindlegenPath = async () => {
    await setKindlegenValue("");
  };
</script>

<ContentLayout {hide}>
  <h2>Configure Runik</h2>
  <!-- When we have more than one setting, make this a standalone component -->
  <div class="setting-entry">
    <span>Path to Kindlegen</span>
    <InfoPopover
      >The path to the kindlegen program on your computer. Kindlegen is a
      utility that comes bundled with Kindle Previewer and is necessary for
      generating kindle dictionaries.</InfoPopover
    >
    <div class="flex">
      <input
        type="text"
        id="kindle-gen"
        on:blur={async (e) => {
          if (
            e.currentTarget.value &&
            e.currentTarget.value != config.kindlegenPath
          ) {
            await setKindlegenValue(e.currentTarget.value);
          }
        }}
        value={config?.kindlegenPath}
      />
      <Button small onClick={handleSetKindlegenPathFromFileSelect}
        >Browse...</Button
      >
      <Button small type={"error"} onClick={handleResetKindlegenPath}
        >Reset</Button
      >
    </div>
  </div>
</ContentLayout>

<style>
  h2 {
    margin-bottom: 32px;
  }
  .setting-entry div {
    margin-top: 1rem;
  }
  .flex {
    display: flex;
    align-items: center;
  }
  /* Unify these styles into a component */
  input {
    height: 28px;
    padding: 1px 8px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    font-size: 0.9rem;
  }
</style>
