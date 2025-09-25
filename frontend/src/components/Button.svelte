<script lang="ts">
  type ButtonType = "primary" | "secondary" | "error";

  export let onClick: (...args: any[]) => void;
  export let type: ButtonType = "primary";
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
    background-color: #202020;
    box-sizing: border-box;
    color: white;
    border: 2px #202020 solid;
    border-radius: 8px;
    padding: 0 32px;
    max-width: 256px;
    height: 48px;
    font-size: 1rem;
    cursor: pointer;
    transition: 200ms all ease-in-out;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  button:hover {
    background-color: #2c2c2c;
    border-color: #2c2c2c;
  }
  .secondary {
    background-color: transparent;
    border: 2px black solid;
    color: black;
  }
  .secondary:hover {
    background-color: #20202018;
  }
  .error {
    background-color: #c76767;
    border: 2px #c76767 solid;
    color: white;
  }
  .error:hover {
    background-color: #b35c5c;
    border-color: #b35c5c;
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
    border-color: #5d5d5d;
    background-color: #5d5d5d;
    color: #d3d3d3;
    cursor: default;
  }
  .secondary:disabled {
    border-color: #5d5d5d;
    background-color: #d3d3d3;
    color: #202020;
    cursor: default;
  }
</style>
