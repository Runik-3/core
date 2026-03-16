/**
 * Commonly used modal declarations.
 */

import { modalStore } from "../stores/modal";
import { device } from "../stores/device";
import { get } from "svelte/store";

export const deviceModal = () => {
  const dev = get(device);
  modalStore.set({
    title: `Device: ${dev.name}`,
    modalType: "device",
    cancelFn: () => modalStore.set(null),
    dangerFn: () => {
      device.disconnect();
      modalStore.set(null);
    },
    cancelLabel: "Close",
    dangerLabel: "Disconnect",
  });
};
