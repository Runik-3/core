<script lang="ts">
  import { modalStore } from "../..//stores/modal";
  import Button from "../Button.svelte";

  const { title, description, confirmFn, confirmLabel, cancelFn } = $modalStore;
  const dict = JSON.parse(description);
  let search = "";

  const confirm = () => {
    confirmFn();
    modalStore.set(null);
  };

  $: filteredDefs = [
    ...new Set([
      // We prioritize word matches over matches within the definition.
      ...dict.filter((def: { Word: string; Definition: string }) => {
        if (!search) {
          return true;
        }
        return def.Word.toLowerCase().includes(search.toLowerCase());
      }),
      ...dict.filter((def: { Word: string; Definition: string }) => {
        if (!search) {
          return true;
        }
        return def.Definition.toLowerCase().includes(search.toLowerCase());
      }),
    ]),
  ];
</script>

<div id="modal-container">
  <div id="modal">
    <div id="modal-header">
      <h2>{title}</h2>
      <input type="text" bind:value={search} placeholder="Search definitions" />
    </div>
    <div id="modal-data">
      {#if Object.keys(filteredDefs).length}
        <table>
          <tbody>
            {#each filteredDefs as def}
              <tr>
                <th>{def.Word}</th>
                <td>{def.Definition}</td>
              </tr>
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
        type="secondary">Cancel</Button
      >
      <div id="btn-divider"></div>
      <Button onClick={confirm} maxWidth small
        >{confirmLabel || "Confirm"}</Button
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
  table * {
    padding: 12px 4px;
  }
  table {
    border-collapse: collapse;
  }
  tr th {
    text-align: left;
    min-width: 144px;
    vertical-align: top;
  }
</style>
