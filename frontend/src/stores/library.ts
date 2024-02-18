import type { DictFile } from "src/types/dict";
import type { Response } from "src/types/response";
import { writable } from "svelte/store";
import { GetDictFiles } from "../../wailsjs/go/main/App"

const { subscribe, set } = writable<DictFile[]>([])

function createLibraryStore() {
  return {
    subscribe,
    fetchDicts: async () => {
      const res: Response<DictFile[]> = await GetDictFiles()
      if (res.Error) {
        throw new Error(res.Error)
      }

      const dictFiles = res.Data.filter((dict) => dict.Extension === "json")
      set(dictFiles)
    }
  }
}

export const library = createLibraryStore()
