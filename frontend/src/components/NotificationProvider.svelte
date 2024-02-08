<script lang="ts">
  import { notifications } from "../stores/notification";
  import Notification from "./Notification.svelte";
</script>

<div id="notification-container">
  {#each $notifications as notification, i}
    <Notification {notification} key={i} />
  {/each}
</div>

<style>
  #notification-container {
    position: absolute;
    bottom: 24px;
    right: 24px;
  }
</style>
