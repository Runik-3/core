import { writable } from "svelte/store";

type ModalType = "convertSelect" | "renameDict" | "device" | "details" | "basic"

export interface ModalProps {
  title: string
  description?: string
  modalType?: ModalType
  confirmFn?: () => void
  cancelFn?: () => void
  dangerFn?: () => void
  confirmLabel?: string
  cancelLabel?: string
  dangerLabel?: string 
  // Specific to user input modals
  formData?: string | null
}

export const modalStore = writable<ModalProps | null>(null);
