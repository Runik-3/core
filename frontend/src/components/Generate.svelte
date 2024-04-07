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

  let dictName = "";
  let depth = 2;
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
    getNameFromUrl(url);
  };

  const getNameFromUrl = (url: string) => {
    dictName = new URL(url).hostname.split(".")[0];
  };

  // TODO: This is hacky and probably should be built into builder
  const buildLanguageUrl = (url: string, languageCode: string) => {
    return `${new URL(url).origin}/${languageCode}/api.php`;
  };

  const resetPageState = () => {
    url = "";
    wikiInfo = null;
  };

  const validDepth = () => {
    const MAX_DEPTH = 5;
    const MIN_DEPTH = 1;

    if (depth > MAX_DEPTH) {
      return MAX_DEPTH;
    }
    if (depth < MIN_DEPTH) {
      return MIN_DEPTH;
    }
    return depth;
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const { Error } = await BuildDictionary(
      wikiUrl,
      dictName,
      validDepth(),
      "json",
    );
    loading = false;
    if (Error) {
      notifications.addNotification({
        message: Error,
        severity: Severity.error,
      });
      return;
    }
    notifications.addNotification({
      message: `Successfully generated ${dictName}`,
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
  <div class="input-container">
    <label for="wiki-input"><strong>Wiki Url</strong></label>
    <div>
      <input
        id="wiki-input"
        placeholder="https://kingkiller.fandom.com/wiki/Kingkiller_Wiki"
        type="text"
        bind:value={url}
      />
      <Button disabled={loading} onClick={() => wikiDetails(url)} small
        >Find</Button
      >
    </div>
  </div>

  <div class="input-container depth-container">
    <label for="depth-input"><strong>Depth</strong></label>
    <input id="depth-input" type="number" bind:value={depth} />
  </div>

  {#if loading}
    <div>
      <Loader />
    </div>
  {:else if wikiInfo?.SiteName}
    <h3>
      Generate dictionary from <span>{wikiInfo?.SiteName}</span>?
    </h3>
    <div class="info-block">
      <p class="info-line">
        <strong>Name:</strong>
        <input id="name-input" type="text" bind:value={dictName} />
      </p>
      <p class="info-line"><strong>Entries:</strong> {wikiInfo?.Articles}</p>
      <p class="info-line"><strong>Language:</strong> {wikiInfo?.Lang}</p>
    </div>
    <Button small disabled={loading} onClick={() => buildDict(url)}
      >Generate</Button
    >
    <div class="info-block">
      {#if wikiInfo.Languages.length > 0}
        <details>
          <summary>Other Available Languages:</summary>
          {#each wikiInfo.Languages as lang}
            <div class="supported-lang">
              <p>
                {lang.autonym || lang.lang}
              </p>
              <Button
                small
                onClick={() => buildDict(buildLanguageUrl(lang.url, lang.lang))}
                >Generate</Button
              >
            </div>
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
    margin-top: 32px;
  }
  h3 > span {
    color: #1f797e;
    font-family: "source serif";
  }
  .info-block {
    margin: 24px 0;
  }
  .info-line {
    display: flex;
    align-items: center;
    height: 36px;
  }
  .info-line strong {
    margin-right: 4px;
  }
  .supported-lang {
    display: flex;
    justify-content: space-between;
    height: 40px;
    margin: 0 8px;
  }
  input {
    padding: 8px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    font-size: 0.9rem;
  }
  #wiki-input {
    min-width: 80%;
  }
  .input-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .input-container input {
    margin-top: 8px;
  }
  .depth-container {
    margin-top: 16px;
  }
  #depth-input {
    width: 36px;
  }
  #name-input {
    width: 112px;
  }
  details {
    width: 256px;
    overflow-y: auto;
    padding: 16px;
    border-radius: 8px;
    border: 1px solid lightgrey;
  }
  details[open] {
    max-height: 270px;
  }
  summary {
    cursor: pointer;
  }
  details[open] summary {
    margin-bottom: 16px;
  }
</style>
