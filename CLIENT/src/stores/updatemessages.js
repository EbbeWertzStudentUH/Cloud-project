import { writable } from "svelte/store";

export const messages = writable([]);

export function addMessage(message) {
    messages.update((m) => [...m, message])
}

export function onMessageType(type, callback) {
    messages.subscribe((mess) => {
        const index = mess.findIndex(message => message.type === type);
        if (index !== -1) {
          const [message] = mess.splice(index, 1); // remove message
            callback(message.data);
            messages.set(mess);
        }
      });
  }

  export function onUpdateMessageType(type, callback) {
    messages.subscribe((mess) => {
        const index = mess.findIndex(message => (message.type === "update" && message.data.type === type));
        if (index !== -1) {
          const [message] = mess.splice(index, 1); // remove message
            callback(message.data.subject, message.data.data);
            messages.set(mess);
        }
      });
  }