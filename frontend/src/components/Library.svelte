<script lang="ts">
  import { Severity, notifications } from "../stores/notification";
  import { library } from "../stores/library";
  import {
    ConvertKoboDictionary,
    DeleteLocalDictFile,
  } from "../../wailsjs/go/main/App";
  import DictListItem from "./DictListItem.svelte";
  import Button from "./Button.svelte";
  import ContentLayout from "./ContentLayout.svelte";
  import { onMount } from "svelte";
  import { device } from "../stores/device";
  import type { Response } from "../types/response";
  import Anvil from "./icons/Anvil.svelte";
  import { nav } from "../stores/nav";

  export let hide = false;

  onMount(async () => {
    try {
      await library.fetchDicts();
    } catch (e) {
      if (e instanceof Error) {
        notifications.addNotification({
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

  const convertDictionary = async (dictName: string) => {
    const res: Response<string> = await ConvertKoboDictionary(dictName);
    return new Promise((resolve, reject) => {
      if (res.Error) {
        return reject(res.Error);
      }
      return resolve(res.Data);
    });
  };

  const sendDictsToDevice = async () => {
    if (!$device?.path) {
      notifications.addNotification({
        message: "No device selected.",
        severity: Severity.info,
        timeout: 5000,
      });
      return;
    }
    try {
      await Promise.all(
        [...selected].map(async (dictName: string) => {
          return convertDictionary(dictName);
        }),
      );
      notifications.addNotification({
        message: "Added to device",
        severity: Severity.success,
        timeout: 5000,
      });
      // refetch dicts and selections
      await device.fetchDicts();
      selected = new Set();
    } catch (e) {
      notifications.addNotification({
        message: `${
          selected.size === 1 ? "Dictionary" : "Dictionaries"
        } failed to send to device:\n ${e}`,
        severity: Severity.error,
      });
    }
  };
</script>

<ContentLayout split {hide}>
  <div>
    <h2>Dictionaries</h2>
    {#if $library.length > 0}
      {#each $library as dict}
        <DictListItem
          {dict}
          {select}
          selected={selected.has(dict.Name)}
          deleteDict={DeleteLocalDictFile}
        />
      {/each}
    {:else}
      <!-- Empty library -->
      <p>
        Nothing here yet. Generate dictionaries in the <a
          id="forge-link"
          href="#"
          on:click={() => nav.set("gen")}>forge</a
        ><Anvil size="36" />
      </p>
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
  #forge-link {
    color: #1f797e;
    text-decoration: none;
  }
</style>
