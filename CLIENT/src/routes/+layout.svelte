<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import '../app.css';

	export let data;

	let currentUser;
	$: data.user.subscribe((v) => {
		currentUser = v;
		console.log(currentUser);
	});

	onMount(async () => {
		await data.fetchUser();
		await data.getFriendsListStuff();
	});

	function acceptFriendRequest(requestId) {
	}

	function copyInviteLink() {
		navigator.clipboard.writeText(`${window.location.origin}/invite/${currentUser.id}`).then(() => {
			alert('Invite link gekopieerd!');
		});
	}
</script>

<div class="flex h-screen flex-col bg-slate-950">
	<nav
		class="flex items-center justify-between border-b-2 border-slate-500 bg-slate-900 p-4 text-white"
	>
		<div class="text-xl font-bold">ProjectCloud</div>
		<div class="text-l">Organise your projects on the cloud</div>
		<div>
			{#if currentUser.id}
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

		<div class="flex-1 ">
			<slot></slot>
		</div>
		<aside class="w-64 bg-slate-900 border-r-2 border-slate-500 text-white flex flex-col">
			<div class="p-4 border-b-2 border-slate-500">
				<h2 class="text-lg font-bold">Friend Requests ({currentUser.friend_requests.length})</h2>
				<ul>
					{#each currentUser.friend_requests as request}
						<li class="flex justify-between items-center mt-2">
							<span>{request.first_name} {request.last_name}</span>
							<button
								on:click={() => acceptFriendRequest(request.id)}
								class="bg-emerald-600 text-white px-2 py-1 rounded"
							>
								Accept
							</button>
						</li>
					{/each}
				</ul>
			</div>

			<div class="p-4 flex-1 overflow-y-auto">
				<h2 class="text-lg font-bold">Friends</h2>
				<ul>
					{#each currentUser.friends as friend}
						<li class="mt-2 flex justify-between">
							<span class="{friend.online ? 'text-emerald-500' : 'text-slate-500'}">
								{friend.first_name} {friend.last_name}
							</span>
							<button on:click={() => data.removeFriend(friend.id)} class="text-red-800">X</button>
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

		
	</div>
</div>

