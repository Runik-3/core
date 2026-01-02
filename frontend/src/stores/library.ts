import type { DictFile } from "../types/dict";
import type { Response } from "../types/response";
import { writable } from "svelte/store";
import { GetLocalDictionaries } from "../../wailsjs/go/main/App"

const { subscribe, set } = writable<DictFile[]>([])

function createLibraryStore() {
  return {
    subscribe,
    fetchDicts: async () => {
      const res: Response<DictFile[]> = await GetLocalDictionaries()
      if (res.Error) {
        throw new Error(res.Error)
      }

      const dictFiles = res.Data
        .sort((a, b) => b.Modified.localeCompare(a.Modified))
      set(dictFiles)
    },
  }
}

export const library = createLibraryStore()

