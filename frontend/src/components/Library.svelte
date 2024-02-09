<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import {
    ConvertKoboDictionary,
    GetDictFiles,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";

  let getDicts = GetDictFiles();
  let selected = new Set();

  const select = (name: string) => {
    if (selected.has(name)) {
      selected.delete(name);
      selected = new Set([...selected]);
      return;
    }
    selected = new Set([...selected, name]);
  };

  const sendDictsToDevice = () => {
    if (selected.size === 0) {
      notifications.addNotificaton({
        message: "You must select a target device from the panel on the right.",
        severity: Severity.info,
        timeout: 5000,
      });
    }
    let error = false;
    selected.forEach(async (dict: string) => {
      // handle error and break out of loop
      const res = await ConvertKoboDictionary(dict);
      error = !!res.Error;
    });
    if (error) {
      notifications.addNotificaton({
        message: `${selected.size === 1 ? "Dictionary" : "Dictionaries"} failed to send to device:\n ${Error}`,
        severity: Severity.error,
      });
    }
  };
</script>

<div id="container">
  <h3>Dictionaries</h3>
  {[...selected]}
  {#await getDicts then dicts}
    {#each dicts as dict}
      {#if dict.Extension === "json"}
        <DictListItem {dict} {select} selected={selected.has(dict.Name)} />
      {/if}
    {/each}
  {/await}
  <button id="send-to-device" on:click={sendDictsToDevice}>
    send to Device
  </button>
</div>

<style>
  #container {
    padding: 24px;
    height: 100%;
    background-color: white;
    border-top-right-radius: 16px;
    box-sizing: border-box;
  }
  h3 {
    padding-bottom: 24px;
  }
</style>
