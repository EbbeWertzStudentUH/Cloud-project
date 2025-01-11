import { GETwithToken, POST, POSTWithToken, DELETEwithToken } from '$lib';
import { writable } from 'svelte/store';
import { goto } from '$app/navigation';

const user = writable({
	id: null,
	first_name: '',
	last_name: '',
	friends: [],
	friend_requests: []
});

export async function load() {
	return { user, fetchUser, logout, removeFriend, getFriendsListStuff, acceptFriendRequest, rejectFriendRequest };
}

async function fetchUser() {
	const token = localStorage.getItem('authToken');
	if (!token) {
		console.log('token is not present, going to login page');
		goto('/login');
	} else {
		console.log(`validating token ${token}...`);
		let data = await GETwithToken('user/authenticate');
		if (data.valid) {
			console.log('token valid!');
			user.update((u) => {
				return {
					...u,
					id: data.user.id,
					first_name: data.user.first_name,
					last_name: data.user.last_name
				};
			});
		} else {
			console.log('token did not validate, going to login page');
			localStorage.removeItem('authToken');
			goto('/login');
		}
	}
}

function logout() {
	localStorage.removeItem('authToken');
	user.set({
		id: null,
		first_name: '',
		last_name: '',
		friends: [],
		friend_requests: []
	});
	goto('/login');
}

async function removeFriend(friend_id) {
	let friends_resp = await DELETEwithToken({friend_id:friend_id},'user/friends');
	let friends = friends_resp.users;
	friends.forEach((f) => {
		f.online = false;
	});
	user.update((u) => {
		return { ...u, friends: friends };
	});
}

async function getFriendsListStuff() {
	let friend_requests_resp = await GETwithToken('user/friend-requests');
	let friends_resp = await GETwithToken('user/friends');
	let friend_requests = friend_requests_resp.users;
	let friends = friends_resp.users;
	friends.forEach((f) => {
		f.online = false;
	});
	user.update((u) => {
		return { ...u, friends: friends, friend_requests: friend_requests };
	});
}

async function acceptFriendRequest(friend_id){
	let friends_resp = await POSTWithToken({friend_id:friend_id},'/user/friend-requests/accept');
	let friends = friends_resp.users;
	friends.forEach((f) => {
		f.online = false;
	});
	user.update((u) => {
		return { ...u, friends: friends };
	});
}
async function rejectFriendRequest(friend_id){
	let friends_requests_resp = await DELETEwithToken({friend_id:friend_id},'/user/friend-requests/reject');
	let friend_requests = friends_requests_resp.users;
	user.update((u) => {
		return { ...u, friend_requests: friend_requests };
	});
}