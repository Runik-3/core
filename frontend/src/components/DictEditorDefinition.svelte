<script lang="ts">
  import Edit from "./icons/Edit.svelte";
  import Check from "./icons/Check.svelte";
  import type { EditableDefinition } from "../types/dict";
  import Garbage from "./icons/Garbage.svelte";

  export let def: EditableDefinition;
  export let anyDefsChanged: boolean;
  export let deleteEntry: (w: string) => void;

  let editMode = false;
  let synonymInput = def.Synonyms ? def.Synonyms.join("\n") : "";

  // Offset by 2px to remove scrollbar
  const scaleTextareaHeight = (element: HTMLElement) => {
    element.style.height = `${element.scrollHeight + 2}px`;
  };

  $: defChanged =
    def.initWord !== def.Word ||
    def.initDefinition !== def.Definition ||
    def.initSynonyms?.join("|") !== def.Synonyms?.join("|");

  $: if (defChanged) {
    anyDefsChanged = true;
  } else {
    anyDefsChanged = false;
  }

  $: def.Synonyms = synonymInput
    .split("\n")
    .map((syn) => syn.trim())
    .filter((syn) => syn);
</script>

<div class="row">
  {#if editMode}
    <textarea
      use:scaleTextareaHeight
      on:input={(e) => scaleTextareaHeight(e.currentTarget)}
      class="editable word"
      bind:value={def.Word}
      placeholder="Word"
      rows="1"
    />
    <textarea
      use:scaleTextareaHeight
      on:input={(e) => scaleTextareaHeight(e.currentTarget)}
      class="editable synonyms"
      bind:value={synonymInput}
      placeholder="Synonyms: one per line"
    />
    <textarea
      use:scaleTextareaHeight
      on:input={(e) => scaleTextareaHeight(e.currentTarget)}
      class="editable definition"
      bind:value={def.Definition}
      placeholder="Definition"
    />
    <div class="btn-container">
      <button
        on:click={() => {
          editMode = false;
        }}
        class="edit-btn"><Check size="16px" color="var(--success)" /></button
      >
    </div>
  {:else}
    <span class="editable word">
      {#if defChanged}
        <span class="changed">*</span>
      {/if}
      {def.Word}
    </span>
    <ul class="editable synonyms">
      {#each def.Synonyms || [] as syn}
        <li>{syn}</li>
      {/each}
    </ul>
    <span class="editable definition">{def.Definition}</span>
    <div class="btn-container">
      <button on:click={() => (editMode = true)} class="edit-btn"
        ><Edit color="var(--text-secondary)" /></button
      >
      <button on:click={() => deleteEntry(def.Word)} class="edit-btn"
        ><Garbage color="var(--error)" size="14" /></button
      >
    </div>
  {/if}
</div>

<style>
  .row {
    display: grid;
    grid-template-areas:
      "word definition btns"
      "synonyms definition btns";
    gap: 2px;
    grid-template-rows: max-content 1fr;
    grid-template-columns: minmax(144px, 168px) 3fr 4rem;
    padding: 1rem 0;
    width: 100%;
  }
  .edit-btn {
    background-color: transparent;
    border: None;
    cursor: pointer;
  }
  .editable {
    padding: 8px 8px;
    border: 1px transparent solid;
    font-size: 1rem;
  }
  .word {
    grid-area: word;
    font-weight: bold;
  }
  .synonyms {
    /* Maintain min text height if no synonyms  */
    min-height: 18px;
    grid-area: synonyms;
    white-space: pre-wrap;
    font-weight: normal;
    color: var(--text-secondary);
    font-style: italic;
  }
  ul {
    list-style-position: outside;
    margin-left: 0.5rem;
    list-style-type: "- ";
    font-style: italic;
  }
  textarea.editable {
    outline: none;
    border-radius: 4px;
    border: 1px gray solid;
    background-color: whitesmoke;
    box-sizing: border-box;
    resize: none;
  }
  .word,
  .synonyms {
    max-width: 12rem;
    text-wrap: wrap;
    word-break: break-word;
  }
  .definition {
    grid-area: definition;
    white-space: pre-wrap;
  }
  .btn-container {
    grid-area: btns;
    display: flex;
    flex-direction: column;
    padding: 16px;
  }
  .changed {
    color: grey;
  }
</style>
