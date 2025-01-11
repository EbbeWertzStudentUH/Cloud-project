import { writable } from "svelte/store";

export const friends = writable([]);
export const friend_statuses = writable(new Map());
export const friend_requests = writable([]);

export function setStatus(friend_id, status) {
    friend_statuses.update(statuses => {
      const updatedStatuses = new Map(statuses);
      updatedStatuses[friend_id] = status;
      return updatedStatuses;
    });
  }
function unsetStatus(friend_id) {
    friend_statuses.update(statuses => {
      const updatedStatuses = new Map(statuses);
      updatedStatuses.delete(friend_id);
      return updatedStatuses;
    });
  }

export function updateFriends(newfriends){
    friends.set(newfriends);
}
export function addFriend(friend){
    friends.update(items => [...items, friend]);
}
export function removeFriend(friend){
    friends.update(items =>
        items.filter(f => f.id !== friend.id)
      );
      unsetStatus(friend.id);
}

export function updateFriendRequests(newfriend_requests){
    friend_requests.set(newfriend_requests);
}
export function addFriendRequest(friend_request){
    friend_requests.update(items => [...items, friend_request]);
}
export function removeFriendRequest(friend_request){
    friend_requests.update(items =>
        items.filter(f => f.id !== friend_request.id)
      );
}