import { writable } from "svelte/store";

export enum Severity {
  success = "success",
  info = "info",
  warn = "warn",
  error = "error",
}

export interface Notification {
  message: string;
  severity: Severity;
  /** Timeout in milliseconds. If not defined, notifications can only be
  dismissed by user action. */
  timeout?: number;
}

const { subscribe, update } = writable<Notification[]>([]);

function createNotificationStore() {
  return {
    subscribe,
    addNotificaton: (notify: Notification) =>
      update((msgs) => {
        return [...msgs, notify];
      }),
    dismissNotification: (index: number) => {
      update((msgs) => {
        return msgs.filter((_, idx) => idx !== index);
      });
    },
  };
}

export const notifications = createNotificationStore();
