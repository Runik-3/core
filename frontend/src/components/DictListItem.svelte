<script lang="ts">
  import type { DictFile, Dict } from "../types/dict";
  import { notifications, Severity } from "../stores/notification";
  import Garbage from "./icons/Garbage.svelte";
  import { library } from "../stores/library";
  import type { Response } from "../types/response";
  import { device } from "../stores/device";
  import { modalStore } from "../stores/modal";
  import { ReadLocalDictionary } from "../../wailsjs/go/main/App";
  import Show from "./icons/Show.svelte";

  export let dict: DictFile;
  export let selected = false;
  export let select: (name: string) => void;
  export let deleteDict: (name: string) => Promise<Response<any>>;
  export let compact = false;

  const deleteDictionary = (dict: DictFile) => {
    modalStore.set({
      title: "Confirm delete",
      description: `Would you like to permanently delete the ${dict.Display} dictionary?`,
      confirmFn: () => confirmedDelete(dict),
    });
  };

  const loadDict = async (dict: DictFile) => {
    const res: Response<Dict> = await ReadLocalDictionary(dict.Name);

    modalStore.set({
      title: res.Data.Name,
      description: JSON.stringify(res.Data.Lexicon) || "",
      confirmFn: () => {},
      modalType: "dict",
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
    // TODO: clean this up with some dependency injection -- pass in the right store.
    await library.fetchDicts();
    await device.fetchDicts();
  };
  const formatDictSize = (size: number) => {
    const sizeKb = size / 1000;
    return Math.round(sizeKb * 10) / 10; // round to a single decimal place
  };
</script>

<li class={`${compact && "compact"}`}>
  <div>
    {#if !compact}
      <button
        class={`checkbox ${selected && "selected"}`}
        aria-label={`select ${dict.Name}`}
        on:click={() => select(dict.Name)}
      ></button>
    {/if}
    {dict.Display}
  </div>
  <span class="list-btn-container">
    {#if !compact}
      <span>{dict.Modified.split("T")[0]}</span>
      <span>{formatDictSize(dict.Size)} KB</span>
    {/if}
    {#if !compact}
      <button
        class="list-item-button"
        aria-label={`view ${dict.Name}`}
        type="button"
        on:click={() => loadDict(dict)}
      >
        <Show size="17px" color="#5d5d5d" />
      </button>
    {/if}
    <button
      class="list-item-button"
      aria-label={`delete ${dict.Name}`}
      type="button"
      on:click={() => deleteDictionary(dict)}
    >
      <Garbage size="16px" color="#c76767" />
    </button>
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
    color: #5d5d5d;
  }
  .checkbox {
    border: 1px solid black;
    background-color: white;
    height: 16px;
    width: 16px;
    border-radius: 2px;
    cursor: pointer;
    margin-right: 12px;
  }
  .selected {
    background-color: #1f797e;
  }
  .list-btn-container {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr min-content;
    grid-gap: 12px;
    text-align: right;
    align-items: center;
    margin-left: 24px;
  }
  .list-item-button {
    border: none;
    background-color: transparent;
    cursor: pointer;
  }
  button:nth-last-child(2) {
    grid-column: 3/3;
  }
  button:last-child {
    grid-column: 4/4;
  }
</style>
