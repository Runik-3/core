<script lang="ts">
  import ContentLayout from "./ContentLayout.svelte";
  export let hide = false;
</script>

<ContentLayout {hide}>
  <h2>Configuration</h2>
  <p>Settings coming...</p>
</ContentLayout>

<style>
  h2 {
    margin-bottom: 24px;
  }
</style>
