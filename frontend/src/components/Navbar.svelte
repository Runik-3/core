<script lang="ts">
  import { nav, type Nav } from "../stores/nav";
  import { device } from "../stores/device";
  import { modalStore } from "../stores/modal";
  import Anvil from "./icons/Anvil.svelte";
  import Book from "./icons/Book.svelte";
  import Cog from "./icons/Cog.svelte";
  import { SelectDevice } from "../../wailsjs/go/main/App";
  import type { Device as DeviceType } from "../types/device";
  import type { Response } from "../types/response";
  import { notifications, Severity } from "../stores/notification";

  const navigate = (location: Nav) => {
    nav.set(location);
  };

  const handleDevice = async () => {
    if (!$device.name) {
      const deviceRes: Response<DeviceType> = await SelectDevice();
      if (deviceRes.Error) {
        notifications.addNotification({
          message: deviceRes.Error,
          severity: Severity.warn,
          timeout: 5000,
        });
      }
      await device.selectDevice(deviceRes.Data.Path);
      if (!$device.name) {
        return
      }
    }
    modalStore.set({
      title: `Device: ${$device.name}`,
      modalType: "device",
      cancelFn: () => modalStore.set(null),
      confirmFn: () => {
        device.disconnect()
        modalStore.set(null);
      },
      cancelLabel: "Close",
      confirmLabel: "Disconnect",
    });
  };
</script>

<nav>
  <div id="tabs">
    <button
      class={`${$nav === "lib" ? "active" : ""} tab-btn`}
      on:click={() => navigate("lib")}
      ><Book
        size="20px"
        color={`${$nav === "lib" ? "black" : "var(--text-secondary)"}`}
      /><span>Library</span></button
    >
    <button
      type="button"
      class={`${$nav === "gen" ? "active" : ""} tab-btn`}
      on:click={() => navigate("gen")}
      ><Anvil
        size="20px"
        color={`${$nav === "gen" ? "black" : "var(--text-secondary)"}`}
      /><span>Forge</span></button
    >
    <button
      class={`${$nav === "conf" ? "active" : ""} tab-btn`}
      on:click={() => navigate("conf")}
      ><Cog
        size="20px"
        color={`${$nav === "conf" ? "black" : "var(--text-secondary)"}`}
      /><span>Settings</span></button
    >
  </div>
  <button id="device-btn" on:click={handleDevice}>
    <span>{$device.name ? `Device: ${$device.name}` : "Connect device"}</span>
  </button>
</nav>

<style>
  nav {
    height: 100%;
    padding: 0 0.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  #tabs {
    height: 100%;
    display: flex;
    justify-content: start;
    align-items: center;
  }
  .tab-btn {
    padding: 0 12px;
    margin-right: 0.5rem;
    border-radius: 8px;
    height: 80%;
    background-color: transparent;
    border: 1px solid transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--text-secondary);
  }
  .tab-btn:hover {
    background-color: #b5b5b5;
  }
  .active {
    background-color: var(--bg);
    color: var(--text);
  }
  .active:hover {
    background-color: var(--bg);
  }
  .tab-btn span {
    font-size: 0.9rem;
    margin-left: 8px;
  }
  #device-btn {
    border: none;
    background-color: transparent;
    color: var(--primary);
    font-size: 0.8rem;
    padding: 4px 8px;
    border: 2px solid var(--primary);
    border-radius: 4px;
    cursor: pointer;
  }
</style>
