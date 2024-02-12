<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import {
    ConvertKoboDictionary,
    GetDictFiles,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";

  let getDicts = GetDictFiles();
  let selected = new Set();

  const select = (name: string) => {
    if (selected.has(name)) {
      selected.delete(name);
      selected = new Set([...selected]);
      return;
    }
    selected = new Set([...selected, name]);
  };

  const sendDictsToDevice = () => {
    if (selected.size === 0) {
      notifications.addNotificaton({
        message: "You must select a target device from the panel on the right.",
        severity: Severity.info,
        timeout: 5000,
      });
    }
    // TODO: if device is not selected

    let error = false;
    selected.forEach(async (dict: string) => {
      // handle error and break out of loop
      const res = await ConvertKoboDictionary(dict);
      error = !!res.Error;
    });
    if (error) {
      notifications.addNotificaton({
        message: `${selected.size === 1 ? "Dictionary" : "Dictionaries"} failed to send to device:\n ${Error}`,
        severity: Severity.error,
      });
    }
  };
</script>

<ContentLayout split>
  <div>
    <h3>Dictionaries</h3>
    {#await getDicts then dicts}
      {#each dicts as dict}
        {#if dict.Extension === "json"}
          <DictListItem {dict} {select} selected={selected.has(dict.Name)} />
        {/if}
      {/each}
    {/await}
  </div>
  <div id="button-container">
    <Button disabled={!selected.size} onClick={sendDictsToDevice}>
      Send to device
    </Button>
  </div>
</ContentLayout>

<style>
  h3 {
    padding-bottom: 24px;
  }
  #button-container {
    display: flex;
    justify-content: center;
  }
</style>
