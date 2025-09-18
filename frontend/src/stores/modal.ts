import { writable } from "svelte/store";
import type { Dict } from "../types/dict";

export interface ModalProps {
  title: string
  description?: string
  confirmFn?: () => void
  modalType?: string
  cancelFn?: () => void
  confirmLabel?: string
  // Specific to dict modal types
  dictData?: Dict
  // Specific to select modal type
  selected?: string | null
}

export const modalStore = writable<ModalProps | null>(null);
