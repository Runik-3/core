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
    // generic response type not being inferred in models
    const { Data, Error } = await GetWikiDetails(wikiUrl);
    if (Error != null) {
      alert("There was an error fetching wiki details.");
      return;
    }
    wikiInfo = Data;
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const { Error } = await BuildDictionary(wikiUrl, "", 1, "json");
    loading = false;
    if (Error != null) {
      alert("There was an error building the dictionary.");
    }
  };
</script>

<main>
  <div class="left">
    <h3>Generate New Dictionary</h3>
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

      <button disabled={loading} on:click={() => buildDict(url)}
        >Generate</button
      >
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
  </div>
  <div class="right">
    <button on:click={openDir}>Select Reader Device</button> <br />
    <p>
      {devicePath
        ? `Device Selected: ${devicePath}`
        : "Please choose your e-reader device."}
    </p>
  </div>
</main>

<style>
  main {
    display: flex;
    justify-content: space-around;
  }
  .success {
    border: 2px green solid;
  }
  .error {
    border: 2px red solid;
  }
</style>
