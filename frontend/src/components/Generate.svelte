<script lang="ts">
  import { BuildDictionary, GetWikiDetails } from "../../wailsjs/go/main/App";

  let loading = false;
  let url = "";

  let wikiInfo = undefined;

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

<div>
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

    <button disabled={loading} on:click={() => buildDict(url)}>Generate</button>
    <br />
  {/if}
</div>
