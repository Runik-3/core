<script lang="ts">
  import { type Notification, notifications } from "../stores/notification";

  export let notification: Notification;
  export let key: number;

  if (notification.timeout) {
    setTimeout(
      () => notifications.dismissNotification(key),
      notification.timeout,
    );
  }
</script>

<div class={`notification ${notification?.severity}`}>
  {notification?.message}
  <!-- Replace with close icon -->
  <button
    on:click={() => notifications.dismissNotification(key)}
    class="notification-close-btn">x</button
  >
</div>

<style>
  .notification {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 8px;
    padding: 24px;
    width: 240px;
    color: white;
  }
  .success {
    background-color: #6ab27e;
  }
  .info {
    background-color: #1f797e;
  }
  .warn {
    background-color: #e6c164;
    color: black;
  }
  .error {
    background-color: #c76767;
  }
  .notification-close-btn {
    border: none;
    background-color: transparent;
    font-size: 1rem;
  }
</style>
