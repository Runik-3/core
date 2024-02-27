import type { DictFile } from "../types/dict";
import type { Response } from "../types/response";
import { writable, type Updater } from "svelte/store";
import { GetDeviceDictionaries } from "../../wailsjs/go/main/App"

export interface Device {
  name: string;
  path: string;
  dicts?: DictFile[];
}

const { subscribe, update, set } = writable<Device>({ name: undefined, path: undefined })

function createDeviceStore() {
  return {
    subscribe,
    selectDevice: async (path: string) => {

      set({ name: getDeviceNameFromPath(path), path })
      await fetchDicts()
    },
    fetchDicts,
  }
}

export const device = createDeviceStore()

async function fetchDicts() {
  const res: Response<DictFile[]> = await GetDeviceDictionaries()
  if (res.Error) {
    throw new Error(res.Error)
  }

  update(device => ({ ...device, dicts: res.Data }))
}

const getDeviceNameFromPath = (path: string) => {
  const parts = path.split("/");
  return parts[parts.length - 1];
};
