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
  /** The notificaton links to an exernal webpage. */
  externalLink?: string;
  /** Get's set automatically as a random key. */
  key?: string
}

const { subscribe, update } = writable<Notification[]>([]);

function createNotificationStore() {
  return {
    subscribe,
    addNotification: (notify: Notification) =>
      update((msgs) => {
        // TODO: replace this eventually with a uuid
        const key = Math.floor(Math.random() * 100000)
        notify.key = key.toString()
        return [...msgs, notify];
      }),
    dismissNotification: (key: string) => {
      update((msgs) => {
        return msgs.filter((notify) => notify.key !== key);
      });
    },
  };
}

export const notifications = createNotificationStore();
