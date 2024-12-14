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

  $: filteredDefs = dict.filter((def) => {
    if (!search) {
      return true;
    }
    return (
      def.Word.toLowerCase().includes(search.toLowerCase()) ||
      def.Definition.toLowerCase().includes(search.toLowerCase())
    );
  });
</script>

<div id="modal-container">
  <div id="modal">
    <div id="modal-content">
      <h2>{title}</h2>
      <input type="text" bind:value={search} />
      <div id="modal-data">
        {#each filteredDefs as def}
          <p><strong>{def.Word}:</strong> {def.Definition}</p>
        {/each}
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
  input {
    border: black 1px solid;
  }
  #modal-data {
    margin-top: 16px;
    height: 70vh;
    overflow-y: auto;
  }
  p {
    margin-bottom: 8px;
  }
</style>
