import { writable } from "svelte/store";

export const user = writable(null);

export function updateUser(newUser) {
    user.set(newUser);
}

export function unsetUser() {
    user.set(null);
}