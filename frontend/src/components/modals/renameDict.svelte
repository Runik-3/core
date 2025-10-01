<script lang="ts">
  import { modalStore } from "../../stores/modal";
  import Basic from "./basic.svelte";

  const handleInput = (e: Event) => {
    modalStore.set({
      ...$modalStore,
      formData: (e.currentTarget as HTMLInputElement).value,
    });
  };
</script>

<Basic>
  <input type="text" placeholder="Dictionary name" on:input={handleInput} />
</Basic>

<style>
  input {
    box-sizing: border-box;
    padding: 8px;
    border-radius: 8px;
    border: 1px solid lightgrey;
    font-size: 0.9rem;
    width: 100%;
  }
</style>
