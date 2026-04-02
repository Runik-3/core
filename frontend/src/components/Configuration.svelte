<script lang="ts">
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import InfoPopover from "./InfoPopover.svelte";
  import { GetConfig, SelectFile, SetConfig, GetVersion } from "../../wailsjs/go/main/App";
  import type { Response } from "../types/response";
  import type { Config } from "../types/config";
  import { notifications, Severity } from "../stores/notification";
  import { onMount } from "svelte";
  import { debounce } from "../lib/debounce";

  export let hide = false;
  let config: Config | undefined = undefined;
  let version: string | undefined = undefined;

  let themeInputValue = "";
  // TODO: We should move this up a level, probably create config store
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
    themeInputValue = config.theme;

    const versionRes: Response<string> = await GetVersion();
    version = versionRes.Data
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

  const setConfig = debounce(async (conf: Partial<Config>) => {
    // TODO: loading state while config saves
    config = { ...config, ...conf };
    const setRes: Response<string> = await SetConfig(config);
    if (setRes.Error) {
      notifications.addNotification({
        message: setRes.Error,
        severity: Severity.info,
        timeout: 5000,
      });
      return
    }

    notifications.addNotification({
      message: "Configuration updated",
      severity: Severity.info,
      timeout: 2000,
    });

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
  });

  /** Accept value arg  */
  const handleSetKindlegenPathFromFileSelect = async () => {
    const newPath = await handleSelectFile();
    if (newPath) {
      const newConfig = await setConfig({ kindlegenPath: newPath });
      config = newConfig;
    }
  };

  const handleResetKindlegenPath = async () => {
    await setConfig({ kindlegenPath: "" });
  };

  $: {
    if (themeInputValue === "light") {
      document.body.classList.add("light");
      document.body.classList.remove("dark");
    }
    if (themeInputValue === "dark") {
      document.body.classList.add("dark");
      document.body.classList.remove("light");
    }
    if (themeInputValue === "system") {
      document.body.classList.remove("dark");
      document.body.classList.remove("light");
      if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        document.body.classList.add("dark");
      }
      if (window.matchMedia("(prefers-color-scheme: light)").matches) {
        document.body.classList.add("light");
      }
    }
    if (themeInputValue) {
      setConfig({ theme: themeInputValue });
    }
  }
</script>

<ContentLayout {hide}>
  <h2>Configure Runik</h2>
  <!-- When we have more settings, we should make standalone components -->
  <div class="setting-entry">
    <div class="fieldset-label">Theme</div>
    <fieldset>
      <input
        name="theme"
        value="system"
        type="radio"
        id="theme-system"
        bind:group={themeInputValue}
      />
      <label for="theme-system">System</label>
      <input
        name="theme"
        value="light"
        type="radio"
        id="theme-light"
        bind:group={themeInputValue}
      />
      <label for="theme-light">Light</label>
      <input
        name="theme"
        value="dark"
        type="radio"
        id="theme-dark"
        bind:group={themeInputValue}
      />
      <label for="theme-dark">Dark</label>
    </fieldset>
  </div>
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
            e.currentTarget.value != config.kindlegenPath
          ) {
            await setConfig({ kindlegenPath: e.currentTarget.value });
          }
        }}
        value={config?.kindlegenPath}
      />
      <div id="btn-divider"></div>
      <Button small onClick={handleSetKindlegenPathFromFileSelect}
        >Browse...</Button
      >
      <div id="btn-divider"></div>
      <Button small type={"error"} onClick={handleResetKindlegenPath}
        >Reset</Button
      >
    </div>
  </div>
  <div id="version">
    {version || ""}
  </div>
</ContentLayout>

<style>
  h2 {
    margin-bottom: 32px;
  }
  .setting-entry {
    margin: 2rem 0;
  }
  .setting-entry div {
    margin-top: 1rem;
  }
  .setting-entry div {
    margin-top: 1rem;
  }
  .flex {
    display: flex;
    align-items: center;
  }
  /* Unify these styles into a component */
  #kindle-gen {
    min-width: 300px;
    height: 28px;
    padding: 1px 8px;
    border-radius: 8px;
    border: 1px solid var(--outline);
    font-size: 0.9rem;
  }
  fieldset {
    display: flex;
    align-items: center;
    border: none;
    border-radius: 8px;
    max-width: min-content;
    padding: 0;
    background-color: var(--bg-secondary);
  }
  .fieldset-label {
    margin-bottom: 1rem;
  }
  fieldset label {
    background-color: transparent;
    padding: 8px;
    width: 5rem;
    text-align: center;
    border-radius: 8px;
    border: 1px solid transparent;
    color: var(--text-secondary);
    cursor: pointer;
  }
  fieldset input[type="radio"]:checked+label {
    background-color: var(--active);
    border: 1px solid var(--outline-active);
    color: var(--text);
  }
  fieldset input[type="radio"] {
    -webkit-appearance: none;
    appearance: none;
  }
  /* TODO: Remove instances of this spacer in favour of a button prop*/
  #btn-divider {
    width: 4px;
  }
  #version {
    position: absolute;
    bottom: 1rem;
    right: 1rem;
    font-style: italic;
    color: var(--text-secondary);
    font-size: 0.9rem;
  }
</style>
