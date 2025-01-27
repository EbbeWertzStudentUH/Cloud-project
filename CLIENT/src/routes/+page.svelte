<script>
	import { page } from '$app/stores';
	import { addProject, projects_list } from '../stores/projects';
	import { writable } from 'svelte/store';
	import { fade, slide, fly } from 'svelte/transition';
	import { onUpdateMessageType } from '../stores/updatemessages';
	import { user } from '../stores/user';
	import { POSTWithToken } from '$lib';

	let showForm = writable(false);
	let newProject = { name: '', deadline: '', github: '' };

	$: current_projects_list = $projects_list;

	onUpdateMessageType('project_add', (subject, data) => {
		addProject(data);
	});
	onUpdateMessageType('new_project', (subject, data) => {
		addProject(data);
	});

	async function createProject() {
		showForm.set(false);
		await POSTWithToken(
			{ name: newProject.name, deadline: newProject.deadline, github_repo: newProject.github },
			'/project'
		);
	}
</script>

<main class="flex h-full flex-grow items-center justify-center">
	<div class="w-[36rem] rounded border-2 border-emerald-700 bg-slate-800 p-8 shadow-md">
		<h1 class="mb-6 text-center text-2xl font-bold text-emerald-500">Projects</h1>

		{#if $showForm}
			<div in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
				<div
					class="mb-6 rounded border border-emerald-600 bg-slate-700 p-4"
					in:fade={{ duration: 300 }}
					out:fade={{ duration: 200 }}
				>
					<h2 class="mb-4 text-lg font-semibold text-emerald-400">New Project</h2>
					<form on:submit|preventDefault={createProject} class="space-y-4">
						<label class="block">
							Name:
							<input
								type="text"
								bind:value={newProject.name}
								class="w-full rounded border border-slate-600 bg-slate-800 p-2 text-white focus:border-emerald-500 focus:outline-none"
								required
							/>
						</label>
						<label class="block">
							Deadline:
							<input
								type="date"
								bind:value={newProject.deadline}
								class="w-full rounded border border-slate-600 bg-slate-800 p-2 text-white focus:border-emerald-500 focus:outline-none"
								required
							/>
						</label>
						<label class="block">
							GitHub Repository:
							<input
								type="url"
								bind:value={newProject.github}
								class="w-full rounded border border-slate-600 bg-slate-800 p-2 text-white focus:border-emerald-500 focus:outline-none"
								required
							/>
						</label>
						<div class="flex justify-end space-x-4">
							<button
								type="button"
								class="rounded bg-gray-600 px-4 py-2 text-white hover:bg-gray-700"
								on:click={() => showForm.set(false)}
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
			</div>
		{:else}
			<div in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
				<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
					<div class="mb-6 flex justify-center">
						<button
							class="rounded bg-emerald-600 px-6 py-2 text-white shadow-md transition-all duration-200 hover:bg-emerald-700 hover:shadow-emerald-400"
							on:click={() => showForm.set(true)}
						>
							+ Create New Project
						</button>
					</div>
					<ul class="space-y-4">
						{#each $projects_list as project, index}
							<div in:slide={{ duration: 3000 }} out:slide={{ duration: 3000 }}>
								<li
									in:fly={{
										x: index % 2 ? -100 : 100,
										opacity: 0,
										duration: 3000,
										delay: index * 100
									}}
								>
									<a
										href={`/project/${project.id}`}
										class="block transform rounded bg-slate-900 p-4 text-white shadow-lg transition-all duration-200 hover:scale-105 hover:bg-slate-800 hover:shadow-emerald-500"
									>
										<h2 class="text-lg font-semibold">{project.name}</h2>
									</a>
								</li>
							</div>
						{/each}
					</ul>
				</div>
			</div>
		{/if}
	</div>
</main>
