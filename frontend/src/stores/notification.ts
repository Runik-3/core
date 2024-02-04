import { writable } from "svelte/store";

export enum Severity {
  success = "success",
  info = "info",
  warn = "warn",
  error = "error",
}
interface Notification {
  message: string;
  severity: Severity;
}

const { subscribe, update } = writable<Notification[]>([]);

function createNotificationStore() {
  return {
    subscribe,
    addNotificaton: (message: string, severity: Severity) =>
      update((msgs) => {
        return [...msgs, { message, severity }];
      }),
    dismissNotification: (index: number) => {
      update((msgs) => {
        return msgs.filter((_, idx) => idx !== index);
      });
    },
  };
}

export const notifications = createNotificationStore();
