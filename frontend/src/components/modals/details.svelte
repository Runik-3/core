<script lang="ts">
  import { modalStore } from "../../stores/modal";
  import Basic from "./basic.svelte";
  import type { DictFile } from "../../types/dict";
  import { ReadLocalDictionary } from "../../../wailsjs/go/main/App";
  import Loader from "../Loader.svelte";

  const { formData } = $modalStore;
  const dictFile: DictFile = JSON.parse(formData);
</script>

<Basic>
  <div id="info-container">
    {#await ReadLocalDictionary(dictFile.Name)}
      <div id="loader-container">
        <Loader />
      </div>
    {:then dict}
      <h4>Definitions</h4>
      <span>{dict.Data.Lexicon.length}</span>
      <h4>Language</h4>
      <span>{dict.Data.Lang}</span>
      <h4>Source wiki</h4>
      <span>{new URL(dict.Data.ApiUrl).origin}</span>
      <h4>Last modified</h4>
      <span>{new Date(dictFile.Modified).toDateString()}</span>
    {/await}
  </div>
</Basic>

<style>
  #info-container {
    margin-bottom: 1.5rem;
  }
  h4 {
    margin-top: 1rem;
    margin-bottom: 0.25rem;
  }
  #loader-container {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
