<script lang="ts">
  import type { Response } from "../../types/response";
  import { WriteLocalDictionary } from "../../../wailsjs/go/main/App";
  import { modalStore } from "../../stores/modal";
  import Button from "../Button.svelte";
  import DictModalDefinition from "./dictModalDefinition.svelte";
  import { notifications, Severity } from "../../stores/notification";
  import type { Definition, Dict, EditableDefinition } from "../../types/dict";
  import Plus from "../icons/Plus.svelte";
  import { dict } from "../../../wailsjs/go/models";

  // const { title, dictData, cancelFn } = $modalStore;
  export let title: string;
  export let dictData: Dict;

  let lexicon: EditableDefinition[] = dictData.Lexicon.map(
    (def: Definition) => ({
      initWord: def.Word,
      initDefinition: def.Definition,
      ...def,
    }),
  );

  let search = "";
  let anyDefsChanged = false; // Defs edited
  let dictModified = false; // Defs added or deleted

  // this should write the new dict to disk
  const saveEdits = async () => {
    if (anyDefsChanged || dictModified) {
      const res: Response<string> = await WriteLocalDictionary({
        Name: title,
        Lexicon: lexicon,
        ApiUrl: dictData.ApiUrl,
        Lang: dictData.Lang,
      } as unknown as dict.Dict); // FIXME: wails types are incompatible
      if (res.Error) {
        notifications.addNotification({
          message: "Failed to save changes. Please try again.",
          severity: Severity.error,
        });
      } else {
        notifications.addNotification({
          message: `Saved changes to ${title}.`,
          severity: Severity.info,
          timeout: 5000,
        });
      }

      modalStore.set(null);
    }
  };

  const deleteEntry = (word: string) => {
    lexicon = lexicon.filter((entry) => entry.Word !== word);
    dictModified = true;
  };

  let addMode = false;
  let newWord = "";
  let newDefinition = "";
  const addEntry = () => {
    if (newWord && newDefinition) {
      lexicon.push({
        Word: newWord,
        Definition: newDefinition,
        initWord: newWord,
        initDefinition: newDefinition,
      });

      dictModified = true;
      notifications.addNotification({
        message: `${newWord} successfully added to dictionary!`,
        severity: Severity.success,
        timeout: 5000,
      });

      // Reset add state
      addMode = false;
      newWord = "";
      newDefinition = "";
    } else {
      notifications.addNotification({
        message: "Cannot add blank word or definition.",
        severity: Severity.warn,
        timeout: 5000,
      });
    }
  };

  let timeoutID = null;
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
      currPage = 1;
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

  let currPage = 1;
  let pageSize = 200;
  $: maxPages = Math.floor(filteredDefs.length / pageSize + 1);

  $: pageStart = (currPage - 1) * pageSize;
  $: pageEnd =
    pageStart + pageSize < filteredDefs.length
      ? pageStart + pageSize
      : filteredDefs.length;
  $: page = filteredDefs.slice(pageStart, pageEnd);
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
      <div id="paging-container">
        <button
          on:click={() => {
            if (currPage !== 1) {
              currPage--;
            }
          }}>&lt</button
        >
        <div id="paging-select-container">
          page:
          <select bind:value={currPage}>
            {#each { length: maxPages } as _, i}
              <option value={i + 1}>{i + 1}</option>
            {/each}
          </select>
          <span>/{maxPages}</span>
        </div>
        <button
          on:click={() => {
            if (currPage !== maxPages) {
              currPage++;
            }
          }}>&gt</button
        >
      </div>
    </div>
    <div id="modal-data">
      {#if Object.keys(filteredDefs).length}
        <table>
          <tbody>
            {#each page as def}
              <DictModalDefinition {def} bind:anyDefsChanged {deleteEntry} />
            {/each}
          </tbody>
        </table>
      {:else if search && !Object.keys(page).length}
        <p>No matches for <span>{search}</span></p>
      {:else}
        <p>This dictionary is empty.</p>
      {/if}
    </div>
    <div id="footer">
      <div id="sub-footer">
        <div id="add-definition-container">
          <Button
            onClick={() => (addMode = !addMode)}
            small
            type={addMode ? "error" : "secondary"}
            ><span>{addMode ? "Cancel" : "Add definition"}</span></Button
          >
          <div id="new-definition" class={addMode ? "" : "hide"}>
            <input
              bind:value={newWord}
              placeholder="word"
              class="new-definition-input"
              type="text"
            />
            <input
              bind:value={newDefinition}
              placeholder="definition"
              class="new-definition-input"
              type="text"
            />
            <Button onClick={addEntry} small type="secondary"
              ><span style="margin-right: 8px;">Add</span><Plus
                color="var(--success)"
              /></Button
            >
          </div>
        </div>
        <div id="count">
          showing {pageStart + 1}-{pageEnd} of {lexicon.length} entries
        </div>
      </div>
      <div id="modal-buttons">
        <Button
          onClick={saveEdits}
          maxWidth
          disabled={!anyDefsChanged && !dictModified}
          small>Save changes</Button
        >
      </div>
    </div>
  </div>
</div>

<style>
  /* TODO - refactor this out of the modal, change names etc. */
  #modal-container {
    /* calc based on header + nav */
    height: calc(100vh - 56px);
  }
  #modal {
    display: grid;
    grid-template-rows: min-content 1fr min-content;
    box-sizing: border-box;
    /* A layer lower than notifications. */
    z-index: 100;
    background-color: var(--bg);
    width: 100%;
    height: 100%;
    border-radius: 8px;
    padding: 1rem;
  }
  #modal-data {
    overflow-y: auto;
    margin-bottom: 1rem;
    border: 1px lightgrey solid;
    border-radius: 8px;
    margin-top: 16px;
    overflow-y: auto;
  }
  #modal-data p {
    padding: 8px;
  }
  #modal-data span {
    text-decoration: underline;
  }
  #modal-header {
    display: flex;
    align-items: center;
  }
  #modal-header input {
    padding: 8px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    font-size: 0.9rem;
    width: 100%;
    margin: 0 1rem;
  }
  #paging-container {
    display: flex;
    align-items: center;
    flex-wrap: nowrap;
    font-size: 0.8rem;
  }
  #paging-select-container {
    display: flex;
    align-items: center;
    font-size: 1rem;
    margin: 0 0.5rem;
    height: 100%;
  }
  #paging-select-container select {
    margin: 0 0.2rem;
    font-size: 1rem;
  }
  #paging-container button {
    height: 2rem;
    width: 2rem;
    border: none;
    background: transparent;
    border: 1px lightgrey solid;
    border-radius: 8px;
    cursor: pointer;
  }
  #footer {
    width: 100%;
  }
  #sub-footer {
    display: flex;
    margin-bottom: 1rem;
  }
  .hide {
    display: none !important;
  }
  #add-definition-container {
    width: 100%;
  }
  #new-definition {
    margin-top: 1rem;
    display: flex;
  }
  #new-definition * {
    margin-right: 8px;
  }
  .new-definition-input {
    border: 1px solid lightgrey;
    padding: 8px;
    font-size: 0.8rem;
    border-radius: 8px;
    width: 100%;
  }
  .new-definition-input:first-of-type {
    width: 25%;
  }

  #count {
    text-align: right;
    width: 50%;
    font-size: 0.8rem;
    font-style: italic;
  }
  #modal-buttons {
    display: flex;
  }
  #btn-divider {
    width: 16px;
  }
</style>
