<script lang="ts">
  import { modalStore } from "../../stores/modal";
  import Basic from "./basic.svelte";
</script>

<Basic>
  <select
    on:input={(e) => {
      modalStore.set({ ...$modalStore, formData: e.currentTarget.value });
    }}
  >
    <option selected disabled hidden></option>
    <option value="dicthtml">DictHtml (Kobo)</option>
    <option value="mobi">Mobi (Kindle)</option>
    <option value="stardict">StarDict (KOReader, Boox)</option>
  </select>
</Basic>

<style>
  select {
    font-size: 1rem;
  }
</style>
