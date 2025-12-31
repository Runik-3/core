<script lang="ts">
  import Ellipsis from "./icons/Ellipsis.svelte";

  export let iconColor: string = "var(--text)";

  interface Item {
    label: string;
    action: () => void;
    icon: any;
    iconProps?: { size: string; color: string };
  }

  export let items: Item[];
  let show = false;
  let container: HTMLElement;

  const toggle = () => (show = !show);

  const handleClickOutside = (e: MouseEvent) => {
    // Already closed
    if (!show) {
      return;
    }
    if (!container.contains(e.target as Node)) {
      show = false;
    }
  };
</script>

<svelte:document on:click={handleClickOutside} />
<div bind:this={container} class="dropdown-container">
  <button class="dropdown-button" on:click={toggle}>
    <Ellipsis size="24px" color={iconColor} />
  </button>
  <ul hidden={!show} id="menu-items-container">
    {#each items as item}
      <li>
        <button
          on:click={() => {
            item.action();
            show = false;
          }}
          ><div class="icon-container">
            <svelte:component this={item.icon} {...item.iconProps} />
          </div>
          <span>{item.label}</span></button
        >
      </li>
    {/each}
  </ul>
</div>

<style>
  .dropdown-container {
    position: relative;
  }
  .dropdown-button {
    display: flex;
    border: none;
    background-color: transparent;
    cursor: pointer;
    border-radius: 8px;
  }
  .dropdown-button:hover {
    background: var(--bg-secondary-hover);
  }
  .dropdown-button:active {
    background: var(--bg-secondary);
  }
  #menu-items-container {
    z-index: 99;
    position: absolute;
    min-width: 176px;
    right: 0;
    background-color: var(--bg);
    border-radius: 0.5rem;
    border: 1px var(--bg-secondary) solid;
    text-align: left;
    padding: 0.5rem 0;
  }
  li {
    padding: 0;
    list-style: none;
  }
  li button {
    display: grid;
    grid-template-columns: 20px 1fr;
    font-size: 1rem;
    padding: 0.5rem 1rem;
    border: none;
    background-color: transparent;
    cursor: pointer;
    width: 100%;
    text-align: left;
    align-items: center;
  }
  li button:hover {
    background: var(--bg-secondary-hover);
  }
  .icon-container {
    justify-self: center;
  }
  li button span {
    margin-left: 0.5rem;
  }
</style>
