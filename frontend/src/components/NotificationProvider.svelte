<script lang="ts">
  import { notifications } from '../stores/notification' 
</script>

<div id="notification-container">
  {#each $notifications as notification, i}
    <div class={`notification ${notification?.severity}`}>
      {notification?.message}
      <!-- Replace with close icon -->
      <button on:click={() => notifications.dismissNotification(i)} class="notification-close-btn">x</button>
    </div>
  {/each}
</div>

<style>
  #notification-container{
    position: absolute;
    bottom: 24px;
    right: 24px;
  }
  .notification{
    display: flex;
    justify-content: space-between;
    align-items: center;
    border: 1px solid;
    border-color: black;
    border-radius: 8px;
    padding: 24px;
    width: 240px;
    color: white;
  }
  .success{
    background-color: green;  
  }
  .info {
    background-color: blue;  
  }
  .warn{
    background-color: yellow;  
    color: black;
  }
  .error{
    background-color: red;  
  }
  .notification-close-btn{
    border: none; 
    background-color: transparent;
    font-size: 1rem;
  }
</style>
