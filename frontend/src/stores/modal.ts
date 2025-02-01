import { writable } from "svelte/store";
import type { Response } from "../types/response";
import type { Dict } from "../types/dict";

export interface ModalProps {
  title: string
  description?: string
  dictData?: Dict
  confirmFn?: () => void
  modalType?: string
  cancelFn?: () => void
  confirmLabel?: string
}

export const modalStore = writable<ModalProps | null>(null);
