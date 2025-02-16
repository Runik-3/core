<script lang="ts">
  import type { Response } from "../../types/response";
  import { WriteLocalDictionary } from "../../../wailsjs/go/main/App";
  import { modalStore } from "../../stores/modal";
  import Button from "../Button.svelte";
  import DictModalDefinition from "./dictModalDefinition.svelte";
  import { notifications, Severity } from "../../stores/notification";
  import type { Definition, EditableDefinition } from "../../types/dict";
  import Plus from "../icons/Plus.svelte";

  const { title, dictData, cancelFn } = $modalStore;
  let lexicon: EditableDefinition[] = dictData.Lexicon.map(
    (def: Definition) => ({
      initWord: def.Word,
      initDefinition: def.Definition,
      ...def,
    }),
  );

  let search = "";
  let anyDefsChanged = false;

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

  const deleteEntry = (word: string) => {
    lexicon = lexicon.filter((entry) => entry.Word !== word);
    anyDefsChanged = true;
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

      anyDefsChanged = true;
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
        <p>No matching definitions.</p>
      {:else}
        <p>This dictionary isempty.</p>
      {/if}
    </div>
    <div id="footer">
      <div id="sub-footer">
        <div id="add-definition-container">
          <Button
            onClick={() => (addMode = !addMode)}
            small
            type={addMode ? "error" : "secondary"}
            ><span>{addMode ? "cancel" : "add definition"}</span></Button
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
                color="#6AB27E"
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
          onClick={cancelFn ? cancelFn : () => modalStore.set(null)}
          maxWidth
          small
          type="secondary">{anyDefsChanged ? "Cancel" : "Close"}</Button
        >
        <div id="btn-divider"></div>
        <Button onClick={saveEdits} maxWidth disabled={!anyDefsChanged} small
          >Save changes</Button
        >
      </div>
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
    margin-bottom: 1rem;
    border: 1px lightgrey solid;
    border-radius: 8px;
    margin-top: 16px;
    overflow-y: auto;
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
