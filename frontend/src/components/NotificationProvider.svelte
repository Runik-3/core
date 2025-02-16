<script lang="ts">
  import { notifications } from "../stores/notification";
  import Notification from "./Notification.svelte";
</script>

<div id="notification-container">
  {#each $notifications as notification (notification.key)}
    <Notification {notification} />
  {/each}
</div>

<style>
  #notification-container {
    /* A layer higher than modals. */
    z-index: 101;
    position: absolute;
    bottom: 24px;
    right: 24px;
  }
</style>
