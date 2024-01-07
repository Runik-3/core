<script lang="ts">
  import {BuildDictionary, SelectDirectory} from '../wailsjs/go/main/App'
  
  let url = ''
  let devicePath = ''

  const openDir = async () => {
    devicePath = await SelectDirectory({Title: "test"})
  }

  const buildDict = async (wikiUrl: string) => {
    const dict = await BuildDictionary(wikiUrl, "", "/Users/matteo/Desktop", 10000, 1, "json")    
    console.log(dict)
  }
</script>

<main>
  <button on:click={openDir}>Select Reader Device</button> <br>
  <p>Device Selected: {devicePath}</p>
  <h3>Generate Dictionary</h3>
  <input placeholder="Wiki Url" type="text" bind:value={url}>
  <button on:click={() => buildDict(url)}>Build</button> <br>
</main>
