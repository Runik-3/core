<script lang="ts">
  export let hide = false;
  // Makes the contentlayout component act more like a background, setting
  // bg color transparent and allowing content to determine padding.
  export let transparent = false;
</script>

<div class={`${hide && "hide"} ${transparent && "transparent"}`}>
  <slot />
</div>

<style>
  div {
    padding: 1rem;
    margin: 0 0.5rem 0.5rem;
    border-radius: 8px;
    /* calc based on header + nav */
    height: calc(100vh - 56px);
    overflow: auto;
    background-color: var(--bg);
    box-sizing: border-box;
  }
  .hide {
    display: none;
  }
  .transparent {
    background-color: transparent;
    padding: 0px;
  }
</style>
