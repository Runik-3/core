import { writable } from "svelte/store";

export interface ModalProps {
  title: string
  description: string
  confirmFn: () => void
  cancelFn?: () => void
  confirmLabel?: string
}

export const modalStore = writable<ModalProps | null>(null);
