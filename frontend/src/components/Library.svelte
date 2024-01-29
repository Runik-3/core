<script lang="ts">
  import {
    ConvertKoboDictionary,
    GetRawDicts,
  } from "../../wailsjs/go/main/App";

  let getDicts = GetRawDicts();
</script>

<div>
  <h3>Library:</h3>
  {#await getDicts then dicts}
    {#each dicts as dict}
      {dict.Name}
      <button on:click={() => ConvertKoboDictionary(dict.Name)}>convert</button
      ><br />
    {/each}
  {/await}
</div>

<style>
  div {
    height: 100%;
    background-color: white;
    border-top-right-radius: 20px;
  }
</style>
