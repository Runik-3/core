<script lang="ts">
  import { modalStore } from "../../stores/modal";
  import Basic from "./basic.svelte";
  import type { Dict, DictFile } from "../../types/dict";
  import { ReadLocalDictionary } from "../../../wailsjs/go/main/App";
  import { onMount } from "svelte";

  const { formData } = $modalStore;
  const dictFile: DictFile = JSON.parse(formData);
  let dict: Dict | undefined;

  onMount(async () => {
    dict = (await ReadLocalDictionary(dictFile.Name)).Data;
  });
</script>

<Basic>
  <div id="info-container">
    <h4>Definitions</h4>
    <span>{dict?.Lexicon.length}</span>
    <h4>Language</h4>
    <span>{dict?.Lang}</span>
    <h4>Source wiki</h4>
    <span>{dict?.ApiUrl}</span>
    <h4>Last modified</h4>
    <span>{new Date(dictFile.Modified).toDateString()}</span>
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
</style>
