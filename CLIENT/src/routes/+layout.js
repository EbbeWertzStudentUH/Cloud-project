import { GETwithToken, POST, POSTWithToken, DELETEwithToken, initializeWebSocket, PUTwithTokenNoResult, killWebSocket } from '$lib';
import { goto } from '$app/navigation';
import { unsetUser, updateUser } from '../stores/user';
import { updateFriendRequests, updateFriends } from '../stores/friends';


export async function load() {
	return { loadUser, logout, deleteFriend, loadFullFriendsList, loadFullFriendRequestsList, acceptFriendRequest, rejectFriendRequest, subscribeToFriends, doInitialRequests };
}

async function loadUser() {
	const token = localStorage.getItem('authToken');
	if (!token) {
		console.log('token is not present, going to login page');
		goto('/login');
	} else {
		console.log(`validating token...`);
		let data = await GETwithToken('/user/authenticate');
		if (data.valid) {
			console.log('token valid!');
			const { id, first_name, last_name } = data.user;
			updateUser({ id, first_name, last_name });
		} else {
			console.log('token did not validate, going to login page');
			logout();
		}
	}
}

async function doInitialRequests(){
	initializeWebSocket();
	await loadFullFriendsList();
	await loadFullFriendRequestsList();
	await subscribeToFriends();
}

function logout() {
	localStorage.removeItem('authToken');
	unsetUser();
	killWebSocket();
	goto('/login');
}

async function subscribeToFriends(){
	await PUTwithTokenNoResult('/notifier/subscribe/friends');
}

async function deleteFriend(friend_id) {
	let resp = await DELETEwithToken({friend_id:friend_id},'/user/friends');
	let friends = resp.users;
	updateFriends(friends);
}

async function loadFullFriendsList() {
	let resp = await GETwithToken('/user/friends');
	let friends = resp.users;
	updateFriends(friends);
}
async function loadFullFriendRequestsList() {
	let resp = await GETwithToken('/user/friend-requests');
	let requests = resp.users;
	updateFriendRequests(requests);
}

async function acceptFriendRequest(friend_id){
	let resp = await POSTWithToken({friend_id:friend_id},'/user/friend-requests/accept');
	let requests = resp.users;
	updateFriendRequests(requests);
	await loadFullFriendsList();
}
async function rejectFriendRequest(friend_id){
	let resp = await DELETEwithToken({friend_id:friend_id},'/user/friend-requests/reject');
	let requests = resp.users;
	updateFriendRequests(requests);
}