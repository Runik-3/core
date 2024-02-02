<script lang="ts">
  import {
    ConvertKoboDictionary,
    GetDictFiles,
  } from "../../wailsjs/go/main/App";

  let getDicts = GetDictFiles();
</script>

<div id="container">
  <h3>Dictionaries</h3>
  {#await getDicts then dicts}
    {#each dicts as dict}
      {#if dict.Extension === "json"}
        <div>
          {dict.Display}
          <button on:click={() => ConvertKoboDictionary(dict.Name)}
            >convert</button
          >
        </div>
      {/if}
    {/each}
  {/await}
</div>

<style>
  #container {
    height: 100%;
    background-color: white;
    border-top-right-radius: 16px;
  }
</style>
