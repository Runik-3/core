<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { library } from "../stores/library";
  import {
    InstallDictionaries,
    DeleteLocalDictFile,
    ExportDictionaries,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import { onMount } from "svelte";
  import { device } from "../stores/device";
  import type { Response } from "../types/response";
  import Anvil from "./icons/Anvil.svelte";
  import { nav } from "../stores/nav";
  import { modalStore } from "../stores/modal";

  export let hide = false;

  onMount(async () => {
    try {
      await library.fetchDicts();
    } catch (e) {
      if (e instanceof Error) {
        notifications.addNotification({
          message: `Failed to find existing dictionaries. \n${e.message}`,
          severity: Severity.error,
        });
      }
    }
  });

  let selected: Set<string> = new Set();

  const toggleSelect = (name: string) => {
    if (selected.has(name)) {
      selected.delete(name);
      selected = new Set([...selected]);
      return;
    }
    selected = new Set([...selected, name]);
  };

  const sendDictsToDevice = async () => {
    const dicts = [...selected];
    const res: Response<string> = await InstallDictionaries(dicts);
    if (res.Error) {
      notifications.addNotification({
        message: `${
          selected.size === 1 ? "Dictionary" : "Dictionaries"
        } failed to send to device:\n ${res.Error}`,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `${selected.size === 1 ? "Dictionary" : "Dictionaries"} added to device`,
      severity: Severity.success,
      timeout: 5000,
    });
    // refetch dicts and selections
    await device.fetchDicts();
    selected = new Set();
  };

  const confirmExport = async (dictType: string) => {
    const dicts = [...selected];
    if (!dictType) {
      notifications.addNotification({
        message: "You must select an export format.",
        severity: Severity.info,
        timeout: 5000,
      });
      return;
    }
    const res: Response<string> = await ExportDictionaries(dictType, dicts);
    if (res.Error) {
      notifications.addNotification({
        message: "There was an error exporting.",
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: res.Data, // Data here is a completion message
      timeout: 5000,
      severity: Severity.success,
    });
    selected = new Set();
  };

  const exportDicts = () => {
    modalStore.set({
      title: "Export",
      description: `Select an export format`,
      confirmLabel: "Choose destination...",
      modalType: "convertSelect",
      confirmFn: () => confirmExport($modalStore.formData),
    });
  };
</script>

<ContentLayout split {hide}>
  <div>
    <h2>My Dictionaries</h2>
    {#if $library.length > 0}
      {#each $library as dict}
        <DictListItem
          {dict}
          {toggleSelect}
          selected={selected.has(dict.Name)}
          deleteDict={DeleteLocalDictFile}
        />
      {/each}
    {:else}
      <!-- Empty library -->
      <div id="empty-library-container">
        <p>
          Nothing here yet. Generate dictionaries in the <a
            id="forge-link"
            href="#"
            on:click={() => nav.set("gen")}>forge</a
          >
        </p>
        <Anvil size="32" />
      </div>
    {/if}
  </div>
  <div id="button-container">
    <Button type="secondary" disabled={!selected.size} onClick={exportDicts}
      >Export as...</Button
    >
    <Button disabled={!selected.size} onClick={sendDictsToDevice}>
      Send to device
    </Button>
  </div>
</ContentLayout>

<style>
  h2 {
    padding-bottom: 32px;
  }
  #button-container {
    display: flex;
    justify-content: space-between;
    width: 340px;
    margin: auto;
  }
  #forge-link {
    color: var(--accent);
    text-decoration: none;
    margin-right: 0.25rem;
  }
  #empty-library-container {
    display: flex;
    align-items: center;
  }
</style>
