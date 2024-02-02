<script lang="ts">
  import {
    ConvertKoboDictionary,
    GetDictFiles,
  } from "../../wailsjs/go/main/App";

  let getDicts = GetDictFiles();
</script>

<div>
  <h3>Dictionaries</h3>
  {#await getDicts then dicts}
    {#each dicts as dict}
      {#if dict.Extension=== "json"}
        {dict.Display}
        <button on:click={() => ConvertKoboDictionary(dict.Name)}>send to device</button>
      {/if}
    {/each}
  {/await}
</div>

<style>
  div {
    height: 100%;
    background-color: white;
    border-top-right-radius: 16px;
  }
</style>
