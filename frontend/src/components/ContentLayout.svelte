<script lang="ts">
  export let hide = false;
  export let split = false;
</script>

<div class={`${split && "split"} ${hide && "hide"}`}>
  <slot />
</div>

<style>
  div {
    padding: 24px;
    height: calc(100vh - 48px);
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
