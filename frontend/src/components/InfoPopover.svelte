<script lang="ts">
  import Info from "./icons/Info.svelte";

  let show = false;
</script>

<span
  on:pointerenter={() => (show = true)}
  on:pointerleave={() => (show = false)}
>
  <Info size="16px" color="#5d5d5d" />
  <div class={`info-popover ${show && "show"}`}>
    <slot />
  </div>
</span>

<style>
  span {
    margin-left: 4px;
    position: relative;
  }
  .info-popover {
    position: absolute;
    top: 0;
    left: 16px;
    display: none;
    opacity: 80%;
    background-color: black;
    color: white;
    padding: 12px;
    width: 294px;
    font-size: 0.9rem;
    line-height: 1.2rem;
    border-radius: 8px;
    border: 1px solid #d3d3d3;
  }
  .show {
    display: block;
  }
</style>
