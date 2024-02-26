<script lang="ts">
  type ButtonType = "primary" | "secondary" | "error"

  export let onClick: (...args: any[]) => void;
  export let type: ButtonType = "primary"
  export let disabled = false;
  export let small = false;
  export let maxWidth = false;
</script>

<button
  class={`${small && "small"} ${maxWidth && "max-width"} ${type}`}
  on:click={onClick}
  {disabled}
>
  <slot />
</button>

<style>
  button {
    background-color: black;
    box-sizing: border-box;
    color: white;
    border: 2px black solid;
    border-radius: 8px;
    padding: 0 32px;
    max-width: 256px;
    height: 48px;
    font-size: 1rem;
    cursor: pointer;
  }
  .secondary {
    background-color: transparent; 
    border: 2px black solid;
    color: black;
  }
  .error {
    background-color: #c76767; 
    border: 2px #c76767 solid;
    color: white;
  }
  .small {
    padding: 0 24px;
    max-width: 256px;
    height: 32px;
    font-size: 0.8rem;
  }
  .max-width {
    max-width: unset;
    width: 100%;
  }
  button:disabled {
    background-color: #5d5d5d;
    color: #d3d3d3;
    cursor: default;
  }
</style>
