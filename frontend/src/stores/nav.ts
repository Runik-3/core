import { writable } from "svelte/store";

export type Nav = "lib" | "gen";
export const nav = writable<Nav>("lib");
