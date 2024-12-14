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
    <div id="modal-content">
      <div id="modal-header">
        <h2>{title}</h2>
        <input
          type="text"
          bind:value={search}
          placeholder="Search definitions"
        />
      </div>
      <div id="modal-data">
        {#if Object.keys(filteredDefs).length}
          <table>
            <thead>
              <th>Word</th>
              <th>Definition</th>
            </thead>
            {#each filteredDefs as def}
              <tr>
                <th>{def.Word}</th>
                <td>{def.Definition}</td>
              </tr>
            {/each}
          </table>
        {:else if search && !Object.keys(filteredDefs).length}
          <p>No matching definitions.</p>
        {:else}
          <p>This dictionary is empty.</p>
        {/if}
      </div>
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
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    padding: 24px;
    z-index: 9999;
    background-color: white;
    width: 90%;
    height: 90%;
    min-width: 300px;
    min-height: 200px;
    border-radius: 8px;
    justify-content: space-between;
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
    border: black 1px solid;
    width: 100%;
    margin-left: 32px;
    padding: 8px;
  }
  #modal-data {
    margin-top: 16px;
    /* TODO - fix this hardcoding, it's not scalable. */
    height: 70vh;
    overflow-y: auto;
  }
  table * {
    border: 1px black solid;
    padding: 4px;
  }
  table {
    border-collapse: collapse;
  }
  thead th {
    padding: 16px;
    min-width: 144px;
  }
  tr th {
    text-align: left;
    vertical-align: top;
  }
</style>
