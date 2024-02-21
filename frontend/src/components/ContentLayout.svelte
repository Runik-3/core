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
    height: 100%;
    background-color: white;
    border-top-right-radius: 16px;
    box-sizing: border-box;
  }
  .split {
    display: grid;
    grid-template-rows: auto 36px;
    grid-template-columns: 1;
  }
  .hide {
    display: none;
  }
</style>
