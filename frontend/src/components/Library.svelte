<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { library } from "../stores/library";
  import {
    InstallDictionaries,
    DeleteLocalDictFile,
    ExportDictionaries,
    ReadLocalDictionary,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import { onMount } from "svelte";
  import { device } from "../stores/device";
  import type { Response } from "../types/response";
  import type { Dict, DictFile } from "../types/dict";
  import Anvil from "./icons/Anvil.svelte";
  import { nav } from "../stores/nav";
  import { modalStore } from "../stores/modal";
  import DictView from "./DictEditor.svelte";

  export let hide = false;

  let checked: Set<string> = new Set();
  let selected: DictFile;
  let anyDefsChanged = false;
  let dictModified = false;

  onMount(async () => {
    try {
      await library.fetchDicts();
      selected = $library[0];
    } catch (e) {
      if (e instanceof Error) {
        notifications.addNotification({
          message: `Failed to find existing dictionaries. \n${e.message}`,
          severity: Severity.error,
        });
      }
    }
  });

  const loadDict = async (dict: DictFile) => {
    const res: Response<Dict> = await ReadLocalDictionary(dict.Name);
    return {
      title: res.Data.Name,
      dictData: res.Data,
    };
  };

  const selectDict = (dict: DictFile) => {
    if (anyDefsChanged || dictModified) {
      modalStore.set({
        title: "Unsaved changes",
        description: "You have unsaved changes to this dictionary.",
        confirmLabel: "Discard changes",
        confirmFn: () => {
          anyDefsChanged = false
          dictModified = false
          selected = dict
        }
      });
    } else {
      selected = dict
    }
  }

  const toggleChecked = (name: string) => {
    if (checked.has(name)) {
      checked.delete(name);
      checked = new Set([...checked]);
      return;
    }
    checked = new Set([...checked, name]);
  };

  $: viewingDict = (async () => await loadDict(selected))();

  const reloadSelectedDict = () => {
    // Force a rerender from dict saved to disk
    selected = selected
    // reset state
    dictModified = false
    anyDefsChanged = false
  }

  // TODO: hook this up again
  const sendDictsToDevice = async () => {
    const dicts = [...checked];
    const res: Response<string> = await InstallDictionaries(dicts);
    if (res.Error) {
      notifications.addNotification({
        message: `${
          checked.size === 1 ? "Dictionary" : "Dictionaries"
        } failed to send to device:\n ${res.Error}`,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `${checked.size === 1 ? "Dictionary" : "Dictionaries"} added to device`,
      severity: Severity.success,
      timeout: 5000,
    });
    // refetch dicts and selections
    await device.fetchDicts();
    checked = new Set();
  };

  const confirmExport = async (dictType: string) => {
    const dicts = [...checked];
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
    checked = new Set();
  };

  // TODO: hook this up again
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

<ContentLayout {hide} transparent={true}>
  {#if $library.length > 0}
    <div id="library-container">
      <div id="dictionary-list">
        <h2>My Dictionaries</h2>
        {#each $library as dict}
          <DictListItem
            {dict}
            {toggleChecked}
            selectDict={(dict) => selectDict(dict)}
            selected={selected === dict}
            checked={checked.has(dict.Name)}
            deleteDict={DeleteLocalDictFile}
          />
        {/each}
      </div>
      <div id="dictionary-editor">
        {#await viewingDict then dict}
          <DictView
            bind:anyDefsChanged
            bind:dictModified
            title={dict.title}
            dictData={dict.dictData}
            reloadDict={reloadSelectedDict}
          />
        {/await}
      </div>
    </div>
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
</ContentLayout>

<style>
  h2 {
    padding-bottom: 24px;
  }
  #library-container {
    display: grid;
    grid-template-columns: 300px 1fr;
    height: 100%;
  }
  #dictionary-list {
    background-color: var(--bg);
    margin-right: 8px;
    border-radius: 8px;
    padding: 1rem;
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
