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
    background-color: var(--primary);
    box-sizing: border-box;
    color: var(--text-contrast);
    border: 2px var(--primary) solid;
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
    background-color: var(--primary-hover);
    border-color: var(--primary-hover);
  }
  .secondary {
    background-color: var(--secondary);
    border: 2px black solid;
    color: black;
  }
  .secondary:hover {
    background-color: var(--secondary-hover);
  }
  .error {
    background-color: var(--error);
    border: 2px var(--error) solid;
    color: var(--text-contrast);
  }
  .error:hover {
    background-color: var(--error-hover);
    border-color: var(--error-hover);
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
    border-color: var(--text-secondary);
    background-color: var(--text-secondary);
    color: var(--bg-secondary);
    cursor: default;
  }
  .secondary:disabled {
    border-color: var(--text-secondary);
    background-color: var(--bg-secondary);
    color: var(--text);
    cursor: default;
  }
</style>
