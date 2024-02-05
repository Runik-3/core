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
    border-radius: 8px;
    padding: 24px;
    width: 240px;
    color: white;
  }
  .success{
    background-color: #6AB27E;  
  }
  .info {
    background-color: #1F797E;  
  }
  .warn{
    background-color: #E6C164;  
    color: black;
  }
  .error{
    background-color: #C76767;  
  }
  .notification-close-btn{
    border: none; 
    background-color: transparent;
    font-size: 1rem;
  }
</style>
