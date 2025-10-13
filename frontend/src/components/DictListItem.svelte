<script lang="ts">
  import type { DictFile, Dict } from "../types/dict";
  import { notifications, Severity } from "../stores/notification";
  import Garbage from "./icons/Garbage.svelte";
  import { library } from "../stores/library";
  import type { Response } from "../types/response";
  import { device } from "../stores/device";
  import { modalStore } from "../stores/modal";
  import {
    ReadLocalDictionary,
    RenameLocalDictionary,
  } from "../../wailsjs/go/main/App";
  import Show from "./icons/Show.svelte";
  import Dropdown from "./Dropdown.svelte";
  import Edit from "./icons/Edit.svelte";

  export let selected = false;
  export let dict: DictFile;
  export let toggleSelect: (name: string) => void;
  export let deleteDict: (name: string) => Promise<Response<any>>;
  export let compact = false;

  const refreshLibrary = () => {
    void library.fetchDicts();
    void device.fetchDicts();
  };

  const deleteDictionary = (dict: DictFile) => {
    modalStore.set({
      title: "Confirm delete",
      description: `Would you like to permanently delete the ${dict.Display} dictionary?`,
      confirmFn: () => confirmedDelete(dict),
    });
  };

  const confirmedDelete = async (dict: DictFile) => {
    const { Error } = await deleteDict(dict.Name);
    if (Error) {
      notifications.addNotification({
        message: `Issue deleting file ${dict.Name}\n${Error}`,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `Deleted ${dict.Display}`,
      severity: Severity.info,
      timeout: 5000,
    });
    refreshLibrary();
  };

  const loadDict = async (dict: DictFile) => {
    const res: Response<Dict> = await ReadLocalDictionary(dict.Name);
    modalStore.set({
      title: res.Data.Name,
      dictData: res.Data,
      modalType: "dict",
    });
  };

  const rename = () => {
    modalStore.set({
      title: "Rename dictionary",
      modalType: "renameDict",
      formData: dict.Display,
      confirmLabel: "Save",
      confirmFn: confirmRename,
    });
  };

  const confirmRename = async () => {
    const newName = $modalStore.formData.trim();
    if (!newName) {
      notifications.addNotification({
        message: "Name field empty",
        severity: Severity.info,
        timeout: 5000,
      });
      return;
    }
    if (newName === dict.Display) {
      notifications.addNotification({
        message: "Please enter a new name",
        severity: Severity.info,
        timeout: 5000,
      });
      return;
    }
    const res: Response<string> = await RenameLocalDictionary(
      dict.Name,
      newName,
    );
    if (res.Error) {
      notifications.addNotification({
        message: `There was an issue attempting to rename ${dict.Display}`,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `Dictionary successfully renamed`,
      severity: Severity.success,
      timeout: 5000,
    });
    refreshLibrary();
  };

  const formatDictSize = (size: number) => {
    const sizeKb = size / 1000;
    return Math.round(sizeKb * 10) / 10; // round to a single decimal place
  };
</script>

<li class={`${compact && "compact"}`}>
  <div id="title-checkbox-container">
    {#if !compact}
      <input
        name={`${dict.Name}-check`}
        id={`${dict.Name}-check`}
        type="checkbox"
        bind:checked={selected}
        on:change={() => toggleSelect(dict.Name)}
      />
    {/if}
    <label for={`${dict.Name}-check`}>
      {dict.Display}
    </label>
  </div>
  <span class={`list-btn-container ${compact ? "compact-list" : ""}`}>
    {#if !compact}
      <span>{dict.Modified.split("T")[0]}</span>
      <span>{formatDictSize(dict.Size)} KB</span>

      <Dropdown
        items={[
          {
            icon: Show,
            iconProps: { size: "18px", color: "var(--primary)" },
            label: "View and edit",
            action: () => loadDict(dict),
          },
          {
            icon: Edit,
            iconProps: { size: "18px", color: "var(--primary)" },
            label: "Rename",
            action: () => rename(),
          },
          {
            icon: Garbage,
            iconProps: { size: "16px", color: "var(--error)" },
            label: "Delete",
            action: () => deleteDictionary(dict),
          },
        ]}
      />
    {:else}
      <button class="list-item-button" on:click={() => deleteDictionary(dict)}
        ><Garbage size={"16px"} color={"var(--error)"} /></button
      >
    {/if}
  </span>
</li>

<style>
  li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 0;
    list-style: none;
    white-space: nowrap;
  }
  .compact {
    padding: 12px 0 0 0;
  }
  span {
    color: var(--text-secondary);
  }

  #title-checkbox-container {
    display: flex;
    align-items: center;
    overflow: auto;
  }
  input[type="checkbox"] {
    /* Hide engine-rendered checkbox */
    -webkit-appearance: none;
    appearance: none;
    background-color: #fff;
    margin: 0;

    /* Our custom styles */
    border: 1px solid black;
    background-color: var(--bg);
    height: 16px;
    width: 16px;
    min-width: 16px;
    border-radius: 2px;
  }
  input[type="checkbox"]:checked {
    background-color: var(--accent);
  }
  label {
    padding: 12px;
  }
  .list-btn-container {
    display: grid;
    grid-template-columns: 1fr 1fr min-content;
    grid-gap: 12px;
    text-align: right;
    align-items: center;
    margin-left: 24px;
  }
  .compact-list {
    grid-template-columns: 1fr min-content;
  }
  .list-item-button {
    border: none;
    background-color: transparent;
    cursor: pointer;
  }
</style>
