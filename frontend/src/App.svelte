<script lang="ts">
  import {
    BuildDictionary,
    SelectDevice,
    GetRawDicts,
    GetWikiDetails,
  } from "../wailsjs/go/main/App";

  let loading = false;
  let url = "";
  let devicePath = "";

  let wikiInfo = undefined;
  let getDicts = GetRawDicts();

  const openDir = async () => {
    devicePath = await SelectDevice();
  };

  const wikiDetails = async (wikiUrl: string) => {
    wikiInfo = await GetWikiDetails(wikiUrl);
    console.log(wikiInfo);
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const dict = await BuildDictionary(wikiUrl, "", 1, "json");
    loading = false;
  };
</script>

<main>
  <button on:click={openDir}>Select Reader Device</button> <br />
  <p>Device Selected: {devicePath}</p>

  <h3>Generate Dictionary</h3>
  <input
    class={wikiInfo?.SiteName ? "success" : ""}
    placeholder="Wiki Url"
    type="text"
    bind:value={url}
  />
  <button disabled={loading} on:click={() => wikiDetails(url)}>Build</button>
  <br />
  {#if wikiInfo?.SiteName}
    <p>
      Generate dictionary from {wikiInfo?.SiteName} with {wikiInfo?.Articles} entries?
    </p>

    <button disabled={loading} on:click={() => buildDict(url)}>Generate</button>
    <br />
  {/if}

  <br />
  <br />
  <h3>Library:</h3>
  {#await getDicts then dicts}
    {#each dicts as dict}
      {dict.Name}<br />
    {/each}
  {/await}
</main>

<style>
  .success {
    border: 2px green solid;
  }
  .error {
    border: 2px red solid;
  }
</style>
