<script lang="ts">
  export let onClick: (...args: any[]) => void;
  export let disabled = false;
  export let small = false;
</script>

<button class={`${small && "small"}`} on:click={onClick} {disabled}>
  <slot />
</button>

<style>
  button {
    background-color: black;
    color: white;
    border: none;
    border-radius: 8px;
    padding: 0 32px;
    max-width: 256px;
    height: 48px;
    font-size: 1rem;
    cursor: pointer;
  }
  .small {
    padding: 0 24px;
    max-width: 256px;
    height: 32px;
    font-size: 0.8rem;
  }
  button:disabled {
    background-color: #5d5d5d;
    color: #d3d3d3;
    cursor: default;
  }
</style>
