<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import '../app.css';

	export let data;

	let currentUser;

	$: data.user.subscribe((v) => {
		currentUser = v;
	});

	onMount(async () => {
		await data.fetchUser();
	});
</script>

<div class="flex h-screen flex-col bg-slate-950">
	<nav
		class="flex items-center justify-between border-b-2 border-slate-500 bg-slate-900 p-4 text-white"
	>
		<div class="text-xl font-bold">ProjectCloud</div>
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
	<slot></slot>
</div>
