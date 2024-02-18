<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { library } from "../stores/library";
  import { ConvertKoboDictionary } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    try {
      await library.fetchDicts();
    } catch (e) {
      if (e instanceof Error) {
        notifications.addNotificaton({
          message: `Failed to find existing dictionaries. \n${e.message}`,
          severity: Severity.error,
        });
      }
    }
  });

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
    // TODO: if device is not selected notify
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
    <h2>Dictionaries</h2>
    {#if $library.length > 0}
      {#each $library as dict}
        {#if dict.Extension === "json"}
          <DictListItem {dict} {select} selected={selected.has(dict.Name)} />
        {/if}
      {/each}
      <!-- Empty library -->
    {:else}
      <p>Nothing here yet. Generate dictionaries in the forge.</p>
    {/if}
  </div>
  <div id="button-container">
    <Button disabled={!selected.size} onClick={sendDictsToDevice}>
      Send to device
    </Button>
  </div>
</ContentLayout>

<style>
  h2 {
    padding-bottom: 24px;
  }
  #button-container {
    display: flex;
    justify-content: center;
  }
</style>
