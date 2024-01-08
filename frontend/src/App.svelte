<script lang="ts">
  import { BuildDictionary, SelectDevice } from "../wailsjs/go/main/App";

  let loading = false;
  let url = "";
  let devicePath = "";

  const openDir = async () => {
    devicePath = await SelectDevice();
  };

  const buildDict = async (wikiUrl: string) => {
    loading = true;
    const dict = await BuildDictionary(
      wikiUrl,
      "",
      "/Users/matteo/Desktop",
      10000,
      1,
      "json"
    );
    loading = false;
    console.log(dict);
  };
</script>

<main>
  <button on:click={openDir}>Select Reader Device</button> <br />
  <p>Device Selected: {devicePath}</p>
  <h3>Generate Dictionary</h3>
  <input placeholder="Wiki Url" type="text" bind:value={url} />
  <button on:click={() => buildDict(url)}>Build</button> <br />
</main>
