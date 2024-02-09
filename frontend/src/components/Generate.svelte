<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { BuildDictionary, GetWikiDetails } from "../../wailsjs/go/main/App";

  let loading = false;
  let url = "";

  let wikiInfo = undefined;

  const wikiDetails = async (wikiUrl: string) => {
    // generic response type not being inferred in models
    const { Data, Error } = await GetWikiDetails(wikiUrl);
    console.log(Data, Error);
    if (Error) {
      notifications.addNotificaton({
        message: Error,
        severity: Severity.info,
        timeout: 5000,
      });
    }
    wikiInfo = Data;
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const { Error } = await BuildDictionary(wikiUrl, "", 1, "json");
    loading = false;
    if (Error) {
      notifications.addNotificaton({
        message: Error,
        severity: Severity.error,
      });
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

<style>
  div {
    height: 100%;
    background-color: white;
    border-top-right-radius: 20px;
  }

  h3 {
    padding-bottom: 24px;
  }
</style>
