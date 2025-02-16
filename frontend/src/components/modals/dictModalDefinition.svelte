<script lang="ts">
  import Edit from "../icons/Edit.svelte";
  import Check from "../icons/Check.svelte";
  import type { EditableDefinition } from "../../types/dict";
  import Garbage from "../icons/Garbage.svelte";

  export let def: EditableDefinition;
  export let anyDefsChanged: boolean;
  export let deleteEntry: (w: string) => void;

  let editMode = false;

  $: defChanged =
    def.initWord !== def.Word || def.initDefinition !== def.Definition;

  $: if (defChanged) {
    anyDefsChanged = true;
  } else {
    anyDefsChanged = false;
  }
</script>

<tr class="row">
  {#if editMode}
    <th bind:textContent={def.Word} contenteditable>{def.Word}</th>
    <td bind:textContent={def.Definition} contenteditable>{def.Definition}</td>
    <td>
      <button
        on:click={() => {
          editMode = false;
        }}
        class="edit-btn"><Check size="16px" color="green" /></button
      >
    </td>
  {:else}
    <th>
      {#if defChanged}
        <span class="changed">*</span>
      {/if}
      {def.Word}
    </th>
    <td>{def.Definition}</td>
    <td>
      <button on:click={() => (editMode = true)} class="edit-btn"
        ><Edit color="#5d5d5d" /></button
      >
      <button on:click={() => deleteEntry(def.Word)} class="edit-btn"
        ><Garbage color="#c76767" size="14" /></button
      >
    </td>
  {/if}
</tr>

<style>
  td,
  th {
    /* Because of how table cells stretch, this acts as min-height */
    height: 92px;
    padding: 16px 8px;
    border: 1px transparent solid;
    text-align: left;
    vertical-align: top;
  }
  td button {
    padding: 8px;
  }
  tr [contenteditable] {
    outline: none;
    border-radius: 4px;
    border: 1px gray solid;
    background-color: whitesmoke;
  }
  tr th {
    min-width: 144px;
  }
  tr td:first-of-type {
    width: 100%;
  }
  .edit-btn {
    background-color: transparent;
    border: None;
    cursor: pointer;
  }
  .changed {
    color: grey;
  }
</style>
