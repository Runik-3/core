<script lang="ts">
  import {
    ConvertKoboDictionary,
    GetDictFiles,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";

  let getDicts = GetDictFiles();
  let selected = new Set()

  const select = (name: string) => {
    if (selected.has(name)) {
      selected.delete(name) 
      selected = new Set([...selected]) 
      return 
    }
    selected = new Set([...selected, name]) 
  }

  const sendDictsToDevice = () => {
    if (selected.size === 0) {
      // handle info feedback "No dictionaries selected"
    }
    selected.forEach(async (dict: string) => {
      // handle error and break out of loop
      await ConvertKoboDictionary(dict)
    })
  }
</script>

<div id="container">
  <h3>Dictionaries</h3>
  {[...selected]}
  {#await getDicts then dicts}
    {#each dicts as dict}
      {#if dict.Extension === "json"}
        <DictListItem dict={dict} select={select} selected={selected.has(dict.Name)} />
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
