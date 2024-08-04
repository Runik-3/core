<script lang="ts">
  import {
    type Notification,
    notifications,
    Severity,
  } from "../stores/notification";
  import Close from "./icons/Close.svelte";
  import Info from "./icons/Info.svelte";
  import Warn from "./icons/Warn.svelte";
  import Error from "./icons/Error.svelte";
  import Success from "./icons/Success.svelte";
  import Link from "./icons/Link.svelte";
  import { BrowserOpenURL } from "../../wailsjs/runtime/runtime"

  export let notification: Notification;

  const severityMap = (severity: Severity) => {
    switch (severity) {
      case Severity.warn:
        return { color: "#e6c164", Component: Warn };
      case Severity.info:
        return { color: "#1f797e", Component: Info };
      case Severity.error:
        return { color: "#c76767", Component: Error };
      case Severity.success:
        return { color: "#6ab27e", Component: Success };
    }
  };
  const { color, Component: Icon } = severityMap(notification.severity);

  if (notification.timeout) {
    setTimeout(
      () => notifications.dismissNotification(notification.key),
      notification.timeout,
    );
  }
</script>

<div class={`notification ${notification.severity}`}>
  <div class={`bar ${notification.severity}`}></div>
  <div class="notification-content">
    <div class="icon">
      <Icon size="24px" {color} />
    </div>
    {notification?.message}
    {#if notification?.externalLink}
      <a class="external-link-btn" href="#" on:click={() => BrowserOpenURL(notification.externalLink)}><Link color='#5D5D5D' /></a> 
    {/if}
  </div>
  <button
    on:click={() => notifications.dismissNotification(notification.key)}
    class="notification-close-btn"
    ><Close size="14px" />
  </button>
</div>

<style>
  .notification {
    display: grid;
    grid-template-rows: 1;
    grid-template-columns: 8px auto 32px;
    border-radius: 8px;
    margin-top: 4px;
    width: 420px;
    background-color: white;
    border-width: 1px;
    border-style: solid;
    overflow: hidden;
    transition: ease-in-out 2s all;
  }
  .notification-content {
    display: flex;
    justify-content: start;
    align-items: center;
    padding: 24px 12px;
    word-break: break-word;
  }
  .icon {
    margin-right: 12px;
  }
  .bar {
    height: 100%;
    border: 4px solid;
    box-sizing: border-box;
  }
  .success {
    border-color: #6ab27e;
  }
  .info {
    border-color: #1f797e;
  }
  .warn {
    border-color: #e6c164;
  }
  .error {
    border-color: #c76767;
  }
  .notification-close-btn {
    padding-top: 8px;
    align-self: start;
    border: none;
    background-color: transparent;
    cursor: pointer;
  }
</style>
