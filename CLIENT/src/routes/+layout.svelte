<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import '../app.css';
	import { initializeWebSocket } from '$lib';
	import { user } from '../stores/user';
	import { addFriend, addFriendRequest, friend_requests, friend_statuses, friends, removeFriend, removeFriendRequest, setStatus } from '../stores/friends';
	import Notifications from '../components/Notifications.svelte';
	import { onMessageType, onUpdateMessageType } from '../stores/updatemessages';

	export let data;
	let { loadUser, logout, deleteFriend, loadFullFriendsList, loadFullFriendRequestsList, acceptFriendRequest, rejectFriendRequest, subscribeToFriends, doInitialRequests } = data;
	$:  currentUser = $user;
	$:  currentFriends = $friends;
	$:  currentFriendStatuses = $friend_statuses;
	$:  currentFriendRequests = $friend_requests;

	onMount(async () => {
		await loadUser();
		await doInitialRequests();
    });

	onUpdateMessageType("new_friend_request", (subject, data) => {
		addFriendRequest(data);
	});
	onUpdateMessageType("removed_friend", (subject, data) => {
		removeFriend(data);
	});
	onUpdateMessageType("new_friend", (subject, data) => {
		addFriend(data);
	});
	onMessageType("subscribed_users_status", (data) => {
		if(data.topic === "friends"){
			data.users.forEach(friend => {
				setStatus(friend.id, friend.status);
			});
		}
	});
	

	function copyInviteLink() {
		navigator.clipboard.writeText(`${window.location.origin}/invite/${currentUser.id}`).then(() => {
			alert('Invite link gekopieerd!');});
	}
</script>

<Notifications />

<div class="flex h-full min-h-screen flex-col bg-slate-950">
	<nav
		class="flex items-center justify-between border-b-2 border-slate-500 bg-slate-900 p-4 text-white"
	>
		<a class="text-xl font-bold" href="/">ProjectCloud</a>
		<div class="text-l">Organise your projects on the cloud</div>
		<div>
			{#if currentUser}
				<span class="mr-4">Hello, {currentUser.first_name} {currentUser.last_name}!</span>
				<button on:click={data.logout} class="rounded bg-slate-700 px-4 py-2">Logout</button>
			{:else}
				<button on:click={() => goto('/login')} class="rounded bg-emerald-600 px-4 py-2"
					>Login</button
				>
				<button on:click={() => goto('/register')} class="ml-2 rounded bg-emerald-800 px-4 py-2"
					>Register</button
				>
			{/if}
		</div>
	</nav>

	<div class="flex flex-1">

		<div class="flex-1">
			<slot></slot>
		</div>
		{#if currentUser }
		<aside class="w-64 bg-slate-900 border-r-2 border-slate-500 text-white flex flex-col">
			<div class="p-4 border-b-2 border-slate-500">
				<h2 class="text-lg font-bold">Friend Requests</h2>
				<ul>
					{#each currentFriendRequests as request}
						<li class="flex justify-between items-center mt-2">
							<span>{request.first_name} {request.last_name}</span>
							<button
								on:click={() => acceptFriendRequest(request.id)}
								class="bg-emerald-600 text-white px-2 py-1 rounded"
							>
								Accept
							</button>
							<button on:click={() => rejectFriendRequest(request.id)} class="text-red-800">X</button>
						</li>
					{/each}
				</ul>
			</div>

			<div class="p-4 flex-1 overflow-y-auto">
				<h2 class="text-lg font-bold">Friends</h2>
				<ul>
					{#each currentFriends as friend}
						<li class="mt-2 flex justify-between">
							<span class="{currentFriendStatuses[friend.id] === "online" ? 'text-emerald-500' : 'text-slate-500'}">
								{friend.first_name} {friend.last_name}
							</span>
							<button on:click={() => deleteFriend(friend.id)} class="text-red-800">X</button>
						</li>
					{/each}
				</ul>
			</div>

			<button
				on:click={copyInviteLink}
				class="bg-emerald-600 text-white p-4 w-full mt-auto"
			>
				Copy Invite Link
			</button>
		</aside>
		{/if}
		
	</div>
</div>

