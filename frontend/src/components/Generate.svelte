<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { BuildDictionary, GetWikiDetails } from "../../wailsjs/go/main/App";
  import ContentLayout from "./ContentLayout.svelte";
  import Button from "./Button.svelte";
  import Loader from "./Loader.svelte";
  import { library } from "../stores/library";
  import type { Response } from "../types/response";
  import type { WikiInfo } from "../types/wikiInfo";

  export let hide = false;

  let loading = false;
  let url = "";

  let wikiInfo = null;

  const wikiDetails = async (wikiUrl: string) => {
    // generic response type not being inferred in models
    loading = true;
    const res: Response<WikiInfo> = await GetWikiDetails(wikiUrl);
    if (Error) {
      notifications.addNotification({
        message: res.Error,
        severity: Severity.info,
        timeout: 5000,
      });
    }
    loading = false;
    wikiInfo = res.Data;
  };

  const getNameFromUrl = (url: string) => {
    return new URL(url).hostname.split(".")[0];
  };

  const resetPageState = () => {
    url = "";
    wikiInfo = null;
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const { Error } = await BuildDictionary(wikiUrl, "", 1, "json");
    loading = false;
    if (Error) {
      notifications.addNotification({
        message: Error,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `Successfully generated ${getNameFromUrl(url)}`,
      severity: Severity.success,
      timeout: 5000,
    });
    // update library
    await library.fetchDicts();
    resetPageState();
  };
</script>

<ContentLayout {hide}>
  <h2>Dictionary Forge</h2>
  <input placeholder="Wiki URL" type="text" bind:value={url} />
  <Button disabled={loading} onClick={() => wikiDetails(url)} small>Find</Button
  >
  {#if loading}
    <div>
      <Loader />
    </div>
  {:else if wikiInfo?.SiteName}
    <h3>
      Generate dictionary from <span>{wikiInfo?.SiteName}</span>?
    </h3>
    <p><strong>Name:</strong> {getNameFromUrl(url)}</p>
    <p><strong>Entries:</strong> {wikiInfo?.Articles}</p>
    <p><strong>Language:</strong> {wikiInfo?.Lang}</p>
    <Button small disabled={loading} onClick={() => buildDict(url)}
      >Generate</Button
    >
  {/if}
</ContentLayout>

<style>
  h2 {
    margin-bottom: 24px;
  }
  h3 {
    margin: 32px 0 12px 0;
  }
  h3 > span {
    color: #1f797e;
  }
  p {
    margin: 8px 0 0 8px;
  }
  p:last-of-type {
    margin-bottom: 24px;
  }
  input {
    padding: 9px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    min-width: 256px;
  }
</style>
