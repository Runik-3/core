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

  let wikiInfo: WikiInfo | null = null;

  const wikiDetails = async (wikiUrl: string) => {
    // generic response type not being inferred in models
    loading = true;
    const res: Response<WikiInfo> = await GetWikiDetails(wikiUrl);
    if (res.Error) {
      notifications.addNotification({
        message: res.Error,
        severity: Severity.info,
        timeout: 5000,
      });
    }
    loading = false;
    wikiInfo = res.Data;
    console.log("ooga booga", wikiInfo);
  };

  const getNameFromUrl = (url: string) => {
    return new URL(url).hostname.split(".")[0];
  };

  const buildLanguageUrl = (url: string, languageCode: string) => {
    return `${new URL(url).origin}/${languageCode}/api.php`;
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
    <div class="info-block">
      <p><strong>Name:</strong> {getNameFromUrl(url)}</p>
      <p><strong>Entries:</strong> {wikiInfo?.Articles}</p>
      <p><strong>Language:</strong> {wikiInfo?.Lang}</p>
    </div>
    <Button small disabled={loading} onClick={() => buildDict(url)}
      >Generate</Button
    >
    <div class="info-block">
      {#if wikiInfo.Languages.length > 0}
        <details>
          <summary>Other Available Languages:</summary>
          {#each wikiInfo.Languages as lang}
            <p>
              {lang.autonym || lang.lang}
              <span
                ><Button
                  small
                  onClick={() =>
                    buildDict(buildLanguageUrl(lang.url, lang.lang))}
                  >Generate</Button
                ></span
              >
            </p>
          {/each}
        </details>
      {/if}
    </div>
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
    font-family: "source serif";
  }
  p {
    margin: 8px 0 0 8px;
  }
  .info-block {
    margin: 24px 0;
  }
  input {
    padding: 9px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    min-width: 256px;
  }
</style>
