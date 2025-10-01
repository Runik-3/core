<script lang="ts">
  import { modalStore } from "../..//stores/modal";
  import Button from "../Button.svelte";

  const { title, description, confirmFn, confirmLabel, cancelFn } = $modalStore;

  const confirm = () => {
    confirmFn();
    modalStore.set(null);
  };
</script>

<div id="modal-container">
  <div id="modal">
    <div id="modal-content">
      <h2>{title}</h2>
      <p>{description}</p>
    </div>
    <slot />
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
    width: 40%;
    height: 25%;
    min-width: 300px;
    min-height: 200px;
    max-width: 500px;
    max-height: 350px;
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
  p {
    margin-top: 16px;
  }
</style>
