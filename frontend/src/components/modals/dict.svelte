<script lang="ts">
  import type { Response } from "../../types/response";
  import { WriteLocalDictionary } from "../../../wailsjs/go/main/App";
  import { modalStore } from "../../stores/modal";
  import Button from "../Button.svelte";
  import DictModalDefinition from "./dictModalDefinition.svelte";
  import { notifications, Severity } from "../../stores/notification";
  import type { Definition, EditableDefinition } from "../../types/dict";

  const { title, dictData, cancelFn } = $modalStore;
  const lexicon: EditableDefinition[] = dictData.Lexicon.map(
    (def: Definition) => ({
      initWord: def.Word,
      initDefinition: def.Definition,
      ...def,
    }),
  );

  let search = "";
  let anyDefsChanged = false;
  let timeoutID = null;

  // this should write the new dict to disk
  const saveEdits = async () => {
    if (anyDefsChanged) {
      const res: Response<string> = await WriteLocalDictionary({
        Name: title,
        Lexicon: lexicon,
      });
      if (res.Error) {
        notifications.addNotification({
          message: "Failed to save changes. Please try again.",
          severity: Severity.error,
        });
      } else {
        notifications.addNotification({
          message: "Successfully saved changes.",
          severity: Severity.info,
          timeout: 5000,
        });
      }

      modalStore.set(null);
    }
  };

  // debounce search
  const handleSearch = (
    e: Event & {
      currentTarget: EventTarget & HTMLInputElement;
      target: HTMLInputElement;
    },
  ) => {
    if (timeoutID) {
      clearTimeout(timeoutID);
      timeoutID = null;
    }
    timeoutID = setTimeout(() => {
      search = e.target.value;
    }, 400);
  };

  $: filteredDefs = [
    ...new Set([
      // We prioritize word matches over matches within the definition.
      ...lexicon.filter((def: EditableDefinition) => {
        return def.initWord.toLowerCase().includes(search.toLowerCase());
      }),
      ...lexicon.filter((def: EditableDefinition) => {
        return def.initDefinition.toLowerCase().includes(search.toLowerCase());
      }),
    ]),
  ];
</script>

<div id="modal-container">
  <div id="modal">
    <div id="modal-header">
      <h2>{title}</h2>
      <input
        type="text"
        on:input={handleSearch}
        placeholder="Search definitions"
      />
    </div>
    <div id="modal-data">
      {#if Object.keys(filteredDefs).length}
        <table>
          <tbody>
            {#each filteredDefs as def}
              <DictModalDefinition {def} bind:anyDefsChanged />
            {/each}
          </tbody>
        </table>
      {:else if search && !Object.keys(filteredDefs).length}
        <p>No matching definitions.</p>
      {:else}
        <p>This dictionary is empty.</p>
      {/if}
    </div>
    <div id="modal-buttons">
      <Button
        onClick={cancelFn ? cancelFn : () => modalStore.set(null)}
        maxWidth
        small
        type="secondary">Close</Button
      >
      <div id="btn-divider"></div>
      <Button onClick={saveEdits} maxWidth disabled={!anyDefsChanged} small
        >Save changes</Button
      >
    </div>
  </div>
</div>

<style>
  #modal-container {
    position: absolute;
    top: 0;
    left: 0;
    height: 100vh;
    width: 100vw;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  #modal {
    display: grid;
    grid-template-rows: min-content 1fr min-content;
    box-sizing: border-box;
    padding: 24px;
    z-index: 9999;
    background-color: white;
    width: 90%;
    height: 90%;
    border-radius: 8px;
  }
  #modal-data {
    overflow-y: auto;
    margin-bottom: 32px;
  }
  #modal-buttons {
    width: 100%;
    display: flex;
  }
  #btn-divider {
    width: 16px;
  }
  #modal-header {
    display: flex;
  }
  #modal-header input {
    padding: 8px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    font-size: 0.9rem;
    width: 100%;
    margin-left: 32px;
  }
  #modal-data {
    margin-top: 16px;
    overflow-y: auto;
  }
</style>
