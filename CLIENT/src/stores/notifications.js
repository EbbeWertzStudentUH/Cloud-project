import { writable } from "svelte/store";

export const notifications = writable([]);

export function addNotification(message) {
    notifications.update((m) => [...m, message])
}

export function removeNotification(notification) {
    notifications.update((n) => n.filter((not) => not !== notification));
}
