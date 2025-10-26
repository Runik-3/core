<script lang="ts">
  export let hide = false;
  export let split = false;
</script>

<div class={`${split && "split"} ${hide && "hide"}`}>
  <slot />
</div>

<style>
  div {
    padding: 1rem;
    margin: 0 0.5rem 0.5rem;
    border-radius: 8px;
    height: calc(100vh - 56px);
    overflow: auto;
    background-color: var(--bg);
    box-sizing: border-box;
  }
  .split {
    display: grid;
    grid-template-rows: auto 48px;
    grid-template-columns: 1;
  }
  .hide {
    display: none;
  }
</style>
