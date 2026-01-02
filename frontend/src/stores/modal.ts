import { writable } from "svelte/store";

type ModalType = "convertSelect" | "renameDict" | "device" | "details" | "basic"

export interface ModalProps {
  title: string
  description?: string
  confirmFn?: () => void
  modalType?: ModalType
  cancelFn?: () => void
  confirmLabel?: string
  cancelLabel?: string
  // Specific to user input modals
  formData?: string | null
}

export const modalStore = writable<ModalProps | null>(null);
