<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { BuildDictionary, GetWikiDetails } from "../../wailsjs/go/main/App";
  import ContentLayout from "./ContentLayout.svelte";
  import Button from "./Button.svelte";
  import Loader from "./Loader.svelte";

  let loading = false;
  let url = "";

  let wikiInfo = null;

  const wikiDetails = async (wikiUrl: string) => {
    // generic response type not being inferred in models
    loading = true;
    const { Data, Error } = await GetWikiDetails(wikiUrl);
    if (Error) {
      notifications.addNotificaton({
        message: Error,
        severity: Severity.info,
        timeout: 5000,
      });
    }
    loading = false;
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
      return;
    }
    notifications.addNotificaton({
      message: `Successfully generated ${wikiInfo.SiteName}`,
      severity: Severity.success,
      timeout: 5000,
    });
  };
</script>

<ContentLayout>
  <h3>Generate New Dictionary</h3>
  <input placeholder="Wiki URL" type="text" bind:value={url} />
  <Button disabled={loading} onClick={() => wikiDetails(url)} small>Find</Button
  >
  {#if loading}
    <div>
      <Loader />
    </div>
  {:else if wikiInfo?.SiteName}
    <p>
      Generate dictionary from {wikiInfo?.SiteName} using {wikiInfo?.Articles} article
      entries?
    </p>

    <Button small disabled={loading} onClick={() => buildDict(url)}
      >Generate</Button
    >
    <br />
  {/if}
</ContentLayout>

<style>
  h3 {
    margin-bottom: 24px;
  }
  input {
    padding: 9px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    min-width: 256px;
  }
</style>
