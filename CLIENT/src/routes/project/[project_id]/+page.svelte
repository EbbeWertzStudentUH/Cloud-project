<script>
	import { onMount } from 'svelte';
	import { addUserToOpenProject, open_project } from '../../../stores/projects';
	import { friends } from '../../../stores/friends';
	import { fade, slide, fly } from 'svelte/transition';
	import { onUpdateMessageType } from '../../../stores/updatemessages';
	import GithubStatsTable from '../../../components/GithubStatsTable.svelte';

	export let data;
	let { fetchProject, addUserToProject, getGithubStats } = data;
	let showUserAddDropdown = false;
	let selectedFriend = null;
	let githubStats = null;
	

	$: current_open_project = $open_project;
	$: currentFriends = $friends;
	onMount(async () => {
		await fetchProject();
		if(current_open_project){
			githubStats = await getGithubStats(current_open_project.github_repo);
		}
	});
	onUpdateMessageType('user_add', (subject, data) => {
		addUserToOpenProject(data);
	});
	async function addUser(){
		showUserAddDropdown = false;
		await addUserToProject(selectedFriend.id);
	}


</script>

<main class="p-4">
	{#if current_open_project}
		<nav class="mb-4">
			<a href="/" class="text-emerald-500 hover:underline">Home</a>
			<span> / </span>
			<span class="font-bold text-slate-300">{current_open_project.name}</span>
		</nav>

		<div class="flex gap-4">
			<aside class="w-1/4 rounded-2xl bg-slate-800 p-4 text-slate-200 shadow-lg">
				<h2 class="mb-3 text-lg font-bold">Milestones</h2>
				<ul class="space-y-2">
					{#each current_open_project.milestones as milestone}
						<li
							class="cursor-pointer rounded-lg bg-slate-700 p-3 shadow transition-colors hover:bg-emerald-500 hover:text-slate-900"
						>
							{milestone.name}
						</li>
					{/each}
				</ul>
			</aside>

			<section class="flex-1 rounded-2xl bg-slate-900 p-6 text-slate-200 shadow-lg">
				<div class="flex justify-between">
					<header class="mb-6">
						<h1 class="text-3xl font-bold text-emerald-400">{current_open_project.name}</h1>
						<p class="mb-2 text-slate-400">
							Deadline: {current_open_project.deadline || 'N/A'}
						</p>
						<a
							href={current_open_project.github_repo}
							target="_blank"
							class="text-blue-500 hover:underline"
						>
							{current_open_project.github_repo}
						</a>
					</header>
					<section class="mb-6 rounded-lg bg-slate-800 p-2">
						<h2 class="mb-3 text-xl font-semibold">Users</h2>
						<ul class="space-y-2">
							{#each current_open_project.users as user}
								<li class="border-b-2 border-slate-900 pr-24 text-slate-500">
									{user.first_name}
								</li>
							{/each}
						</ul>
						<div class="mt-4">
							

							{#if showUserAddDropdown}
								<div class="mt-2 rounded-lg bg-slate-700 p-4 shadow-lg" in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
									<h3 class="mb-3 text-lg text-slate-200">Select a Friend</h3>
									<select
										bind:value={selectedFriend}
										class="w-full rounded-lg border-none bg-slate-600 p-2 text-slate-200 focus:outline-none"
									>
										<option disabled selected value={null}>Select a friend...</option>
										{#each currentFriends.filter((friend) => !current_open_project.users.some((user) => user.id === friend.id)) as friend}
											<option value={friend} class="bg-slate-600 text-slate-200">
												{friend.first_name}
											</option>
										{/each}
									</select>

									{#if selectedFriend}
										<div class="mt-3 flex justify-between gap-4" in:slide={{ y: -20, duration: 300 }}
										out:slide={{ y: -20, duration: 300 }}>
											<button
												on:click={addUser}
												class="rounded-lg bg-emerald-500 px-4 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
											>
												Add Friend
											</button>
											<button
												on:click={() => {
													showUserAddDropdown = false;
													selectedFriend = null;
												}}
												class="rounded-lg bg-slate-500 px-4 py-2 text-slate-900 shadow transition-colors hover:bg-red-600"
											>
												Cancel
											</button>
										</div>
									{/if}
								</div>
							{:else}
							<button in:slide={{ y: -20, duration: 300 }}
							out:slide={{ y: -20, duration: 300 }}
								on:click={() => (showUserAddDropdown = !showUserAddDropdown)}
								class="rounded-lg w-full bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
							>
								Add User
							</button>
							{/if}
						</div>
					</section>
				</div>

				<GithubStatsTable githubStats={githubStats}></GithubStatsTable>
				<section class="flex gap-4">
					<button
						class="rounded-2xl bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
					>
						Add Milestone
					</button>
					<button
						class="rounded-2xl bg-blue-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-blue-600"
					>
						Create Task
					</button>
				</section>
			</section>
		</div>
	{/if}
</main>
