import { writable } from "svelte/store";

export type Nav = "lib" | "gen" | "conf";
export const nav = writable<Nav>("lib");
