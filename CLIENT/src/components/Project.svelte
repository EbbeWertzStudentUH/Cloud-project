<script>
	import { POST } from '$lib';
	import { addUserToProject } from '$lib/project_requests';
	import { friends } from '../stores/friends';
	import { addUserToOpenProject } from '../stores/projects';
	import { onUpdateMessageType } from '../stores/updatemessages';
    import { fade, slide, fly } from 'svelte/transition';
	import GithubStatsTable from './GithubStatsTable.svelte';



    let { project, githubStats } = $props();
    


    let showUserAddDropdown = $state(false);
	let selectedFriend = $state(null);

    let currentFriends = $friends;

    onUpdateMessageType('user_add', (subject, data) => {
		addUserToOpenProject(data);
	});
    async function addUser() {
		showUserAddDropdown = false;
        addUserToProject(selectedFriend.id, project.id)
	}
</script>

<div class="flex justify-between">
    <header class="mb-6">
        <h1 class="text-3xl font-bold text-emerald-400">{project.name}</h1>
        <p class="mb-2 text-slate-400">
            Deadline: {project.deadline}
        </p>
        <a
            href={project.github_repo}
            target="_blank"
            class="text-blue-500 hover:underline"
        >
            {project.github_repo}
        </a>
    </header>
    <section class="mb-6 rounded-lg bg-slate-800 p-2">
        <h2 class="mb-3 text-xl font-semibold">Users</h2>
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
                            <button onclick={() => {
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
                    class="w-full rounded-lg bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
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