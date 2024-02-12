<script lang="ts">
  import type { DictFile } from "../types/dict";
  import { notifications, Severity } from "../stores/notification";
  import Garbage from "./icons/Garbage.svelte";

  export let dict: DictFile;
  export let selected = false;
  export let select: (name: string) => void;

  const notify = () =>
    notifications.addNotificaton({
      message: "Hello there this is an error message.",
      severity: Severity.error,
    });
</script>

<li>
  <div>
    <button
      class={`checkbox ${selected && "selected"}`}
      aria-label={`select ${dict.Name}`}
      on:click={() => select(dict.Name)}
    ></button>
    {dict.Display}
  </div>
  <span class="list-btn-container">
    <span>{dict.Modified.split("T")[0]}</span>
    <span>{dict.Size / 1000} kb</span>
    <button class="list-item-delete" type="button" on:click={notify}>
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
    display: flex;
    align-items: center;
  }
  .list-btn-container > * {
    margin-left: 16px;
  }
  .list-item-delete {
    border: none;
    background-color: transparent;
    cursor: pointer;
  }
</style>