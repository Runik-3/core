<script lang="ts">
  import type { DictFile } from "../types/dict";
  import { notifications, Severity } from "../stores/notification";
  import Garbage from "./icons/Garbage.svelte";
  import { library } from "../stores/library";
  import type { Response } from "../types/response";
  import { device } from "../stores/device";
  import { modalStore } from "../stores/modal";
  import { RenameLocalDictionary } from "../../wailsjs/go/main/App";
  import Dropdown from "./Dropdown.svelte";
  import Edit from "./icons/Edit.svelte";
  import Info from "./icons/Info.svelte";

  export let dict: DictFile;
  // Dictionary is selected
  export let selected = false;
  export let selectDict: (dict: DictFile) => void;
  // Checkbox is checked
  export let checked = false;
  export let toggleChecked: (name: string) => void;
  export let deleteDict: (name: string) => Promise<Response<any>>;
  export let isDeviceList = false;
  export let reloadDict: (newName?: string) => void;

  const refreshLibrary = async () => {
    await library.fetchDicts();
    await device.fetchDicts();
  };

  const dictDetails = (dict: DictFile) => {
    modalStore.set({
      title: dict.Display,
      modalType: "details",
      cancelLabel: "Close",
      // Kind of a hacky way to get details to modal
      formData: JSON.stringify(dict),
    });
  };

  const loadDeviceModal = () => {
    modalStore.set({
      title: `Device: ${$device.name}`,
      modalType: "device",
      cancelFn: () => modalStore.set(null),
      confirmFn: () => {
        device.disconnect();
        modalStore.set(null);
      },
      cancelLabel: "Close",
      confirmLabel: "Disconnect",
    });
  };

  const confirmDelete = async (dict: DictFile) => {
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
    await refreshLibrary();

    if (isDeviceList) {
      loadDeviceModal();
    }
  };

  const deleteDictionary = (dict: DictFile) => {
    modalStore.set({
      title: "Confirm delete",
      description: `Would you like to permanently delete the ${dict.Display} dictionary?`,
      confirmFn: () => confirmDelete(dict),
      cancelFn: () => {
        modalStore.set(null);
        if (isDeviceList) {
          loadDeviceModal();
        }
      },
      modalType: "basic",
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
        message: `There was an issue attempting to rename ${dict.Display}:\n${res.Error}`,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `Dictionary successfully renamed`,
      severity: Severity.success,
      timeout: 5000,
    });

    await refreshLibrary();
    reloadDict(newName);
  };

  const formatDictSize = (size: number) => {
    const sizeKb = size / 1000;
    return Math.round(sizeKb * 10) / 10; // round to a single decimal place
  };
</script>

<li class={`${isDeviceList && "device"} ${selected && "selected"}`}>
  <div id="title-checkbox-container">
    {#if !isDeviceList}
      <input
        name={`${dict.Name}-check`}
        id={`${dict.Name}-check`}
        type="checkbox"
        bind:checked
        on:change={() => toggleChecked(dict.Name)}
      />
    {/if}
    {#if !isDeviceList}
      <button class="dict-label" on:click={() => selectDict(dict)}>
        {dict.Display}
      </button>
    {:else}
      <span class="dict-label">
        {dict.Display}
      </span>
    {/if}
  </div>
  <span class={`list-btn-container ${isDeviceList ? "device-list" : ""}`}>
    {#if !isDeviceList}
      <!-- <span>{dict.Modified.split("T")[0]}</span> -->
      <!-- <span>{formatDictSize(dict.Size)} KB</span> -->

      <Dropdown
        items={[
          {
            icon: Info,
            iconProps: { size: "18px", color: "var(--primary)" },
            label: "Details",
            action: () => dictDetails(dict),
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
    padding: 8px;
    list-style: none;
    white-space: nowrap;
  }
  .selected {
    background-color: var(--secondary-hover);
    border-radius: 8px;
  }
  .device {
    padding: 12px 0 0 0;
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
    margin-right: 0.75rem;

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
  .dict-label {
    background: transparent;
    border: none;
    text-align: left;
    font-size: 1rem;
    padding: 12px 0;
    width: 250px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    cursor: pointer;
    color: var(--text);
  }
  .selected .dict-label {
    color: var(--text);
  }
  .device .dict-label {
    width: 180px;
    cursor: unset;
  }
  .list-btn-container {
    text-align: right;
    align-items: center;
    margin-left: 24px;
  }
  .device-list {
    grid-template-columns: 1fr min-content;
  }
  .list-item-button {
    border: none;
    background-color: transparent;
    cursor: pointer;
  }
</style>
