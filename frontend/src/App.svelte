<script lang="ts">
  import { nav } from "./stores/nav";
  import DevicePanel from "./components/DevicePanel.svelte";
  import Generate from "./components/Generate.svelte";
  import Library from "./components/Library.svelte";
  import Navbar from "./components/Navbar.svelte";
  import NotificationProvider from "./components/NotificationProvider.svelte";
  import ModalProvider from "./components/ModalProvider.svelte";
  import { modalStore } from "./stores/modal";
  import { Severity, notifications } from "./stores/notification";
  import { onMount } from "svelte";
  import { CheckForUpdate } from "../wailsjs/go/main/App";

  // on application start
  onMount(async () => {
    // check for updates
    const updateAvailable = await CheckForUpdate();
    if (updateAvailable) {
      notifications.addNotification({
        message: `A new update is available. Visit runik.app to download the latest version.`,
        severity: Severity.info,
        externalLink: "https://runik.app"
      });
    }
  });
</script>

<main>
  <div id="nav">
    <Navbar />
  </div>
  <div id="view">
    <Library hide={$nav !== "lib"} />
    <Generate hide={$nav !== "gen"} />
    <!-- Disabled until we have global config to add. -->
    <!-- <Configuration hide={$nav !== "conf"} /> -->
  </div>
  <div id="dev">
    <DevicePanel />
  </div>
  <NotificationProvider />
  {#if $modalStore}
    <ModalProvider />
  {/if}
</main>

<style>
  main {
    height: 100%;
    display: grid;
    grid-template-rows: 48px auto;
    grid-template-columns: auto minmax(300px, 350px);
    grid-template-areas:
      "nav nav"
      "view dev";
  }
  #nav {
    grid-area: nav;
  }

  #view {
    grid-area: view;
  }

  #dev {
    grid-area: dev;
  }
</style>
