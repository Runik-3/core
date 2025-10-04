<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { library } from "../stores/library";
  import {
    ConvertDictionary,
    DeleteLocalDictFile,
    ExportDictionary,
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

  const convertDictionary = async (dictName: string) => {
    const res: Response<string> = await ConvertDictionary(dictName);
    return new Promise((resolve, reject) => {
      if (res.Error) {
        return reject(res.Error);
      }
      return resolve(res.Data);
    });
  };

  const sendDictsToDevice = async () => {
    if (!$device?.path) {
      notifications.addNotification({
        message: "No device selected.",
        severity: Severity.info,
        timeout: 5000,
      });
      return;
    }
    try {
      await Promise.all(
        [...selected].map(async (dictName: string) => {
          return convertDictionary(dictName);
        }),
      );
      notifications.addNotification({
        message: "Added to device",
        severity: Severity.success,
        timeout: 5000,
      });
      // refetch dicts and selections
      await device.fetchDicts();
      selected = new Set();
    } catch (e) {
      notifications.addNotification({
        message: `${
          selected.size === 1 ? "Dictionary" : "Dictionaries"
        } failed to send to device:\n ${e}`,
        severity: Severity.error,
      });
    }
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
    const res: Response<string> = await ExportDictionary(dictType, dicts);
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
    console.log(selected);
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
      <p>
        Nothing here yet. Generate dictionaries in the <a
          id="forge-link"
          href="#"
          on:click={() => nav.set("gen")}>forge</a
        ><Anvil size="36" />
      </p>
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
  }
</style>
