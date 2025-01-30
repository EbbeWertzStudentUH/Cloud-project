<script>
	import { POST } from '$lib';
	import { addUserToProject, makeMilestoneInProject } from '$lib/project_requests';
	import { friends } from '../stores/friends';
	import { addUserToOpenProject } from '../stores/projects';
	import { fade, slide, fly } from 'svelte/transition';
	import GithubStatsTable from './GithubStatsTable.svelte';

	let { project, githubStats, hasGithub } = $props();

	let showUserAddDropdown = $state(false);
	let selectedFriend = $state(null);
	let showForm = $state(false);
	let newMilestone = $state({ name: '', deadline: '' });

	let currentFriends = $state([]);
	$effect(() => {
        currentFriends = $friends;
    });

	async function addUser() {
		showUserAddDropdown = false;
		addUserToProject(selectedFriend.id, project.id);
	}
	async function createMilestone() {
		showForm = false;
        await makeMilestoneInProject(project.id, newMilestone.name, newMilestone.deadline)
	}
</script>

<div class="flex justify-between gap-8">
	<section class="w-full">
		<header class="mb-6">
			<h1 class="text-3xl font-bold text-emerald-400">{project.name}</h1>
			<p class="mb-2 text-slate-400">
				Deadline: {project.deadline}
			</p>
			<a href={project.github_repo} target="_blank" class="text-blue-500 hover:underline">
				{project.github_repo}
			</a>
		</header>
		{#if showForm}
			<section
				class="mb-6 rounded border border-emerald-600 bg-slate-700 p-4"
				in:slide={{ y: -20, duration: 300 }}
				out:slide={{ y: -20, duration: 300 }}
			>
				<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
					<h2 class="mb-4 text-lg font-semibold text-emerald-400">New Milestone</h2>
					<form onsubmit={createMilestone} class="space-y-4">
						<label class="block">
							Name:
							<input
								type="text"
								bind:value={newMilestone.name}
								class="w-full rounded border border-slate-600 bg-slate-800 p-2 text-white focus:border-emerald-500 focus:outline-none"
								required
							/>
						</label>
						<label class="block">
							Deadline:
							<input
								type="date"
								bind:value={newMilestone.deadline}
								class="w-full rounded border border-slate-600 bg-slate-800 p-2 text-white focus:border-emerald-500 focus:outline-none"
								required
							/>
						</label>
						<div class="flex justify-end space-x-4">
							<button
								type="button"
								class="rounded bg-gray-600 px-4 py-2 text-white hover:bg-gray-700"
								onclick={() => (showForm = false)}
							>
								Cancel
							</button>
							<button
								type="submit"
								class="rounded bg-emerald-600 px-4 py-2 text-white hover:bg-emerald-700"
							>
								Save
							</button>
						</div>
					</form>
				</div>
			</section>
		{:else}
			<section class="mb-6" in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
				<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
					<button
						onclick={() => {
							showForm = true;
						}}
						class="rounded-md bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
					>
						Add Milestone
					</button>
				</div>
			</section>
		{/if}
	</section>

	<section class="mb-6 rounded-lg bg-slate-800 p-2 h-min">
		<h2 class="mb-3 text-xl font-semibold">Team</h2>
		<ul class="space-y-2">
			{#each project.users as user}
				<li class="border-b-2 border-slate-900 pr-24 text-slate-500">
					{user.first_name}
				</li>
			{/each}
		</ul>
		<div class="mt-4">
			{#if showUserAddDropdown}
				<div
					class="mt-2 rounded-lg bg-slate-700 p-4 shadow-lg"
					in:slide={{ y: -20, duration: 300 }}
					out:slide={{ y: -20, duration: 300 }}
				>
					<h3 class="mb-3 text-lg text-slate-200">Select a Friend</h3>
					<select
						bind:value={selectedFriend}
						class="w-full rounded-lg border-none bg-slate-600 p-2 text-slate-200 focus:outline-none"
					>
						<option disabled selected value={null}>Select a friend...</option>
						{#each currentFriends.filter((friend) => !project.users.some((user) => user.id === friend.id)) as friend}
							<option value={friend} class="bg-slate-600 text-slate-200">
								{friend.first_name}
							</option>
						{/each}
					</select>

					{#if selectedFriend}
						<div
							class="mt-3 flex justify-between gap-4"
							in:slide={{ y: -20, duration: 300 }}
							out:slide={{ y: -20, duration: 300 }}
						>
							<button
								onclick={addUser}
								class="rounded-lg bg-emerald-500 px-4 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
							>
								Add Friend
							</button>
							<button
								onclick={() => {
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
				<button
					in:slide={{ y: -20, duration: 300 }}
					out:slide={{ y: -20, duration: 300 }}
					onclick={() => (showUserAddDropdown = !showUserAddDropdown)}
					class="w-full rounded-md bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
				>
					Add Member
				</button>
			{/if}
		</div>
	</section>
</div>
{#if hasGithub}
	<GithubStatsTable {githubStats}></GithubStatsTable>
{/if}
