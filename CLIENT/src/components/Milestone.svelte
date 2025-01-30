<script>
	import { makeTaskInMilestone } from '$lib/project_requests';
	import { addTaskToMilestone, open_project } from '../stores/projects';
	import Task from './Task.svelte';
	import { fade, slide, fly } from 'svelte/transition';

	let { milestone_id } = $props();
	let showForm = $state(false);
	let newTask = $state({ name: '' });

    let milestone = $derived( 
        $open_project.milestones.find(m => m.id === milestone_id) || {}
    );

	async function createTask() {
		showForm = false;
        await makeTaskInMilestone($open_project.id, milestone.id, newTask.name);
	}
</script>

<header class="mb-6">
	<h1 class="text-3xl font-bold text-emerald-400">
		<span class="text-xl font-bold text-slate-500">Milestone: </span>
		{milestone.name}
	</h1>
	<p class="mb-2 text-slate-400">
		Deadline: {milestone.deadline}
	</p>
</header>
<section class="mb-4">
	{#if showForm}
		<section
			class="mb-6 rounded border border-emerald-600 bg-slate-700 p-4"
			in:slide={{ y: -20, duration: 300 }}
			out:slide={{ y: -20, duration: 300 }}
		>
			<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
				<h2 class="mb-4 text-lg font-semibold text-emerald-400">New Task</h2>
				<form onsubmit={createTask} class="space-y-4">
					<label class="block">
						Name:
						<input
							type="text"
							bind:value={newTask.name}
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
		<section in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
			<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
				<button
					onclick={() => {
						showForm = true;
					}}
					class="rounded-md bg-emerald-500 px-6 py-2 text-slate-900 shadow transition-colors hover:bg-emerald-600"
				>
					Add Task
				</button>
			</div>
		</section>
	{/if}
</section>
<section>
	<h2 class="mb-3 text-lg font-bold">Tasks:</h2>
	<ul class="space-y-2">
		{#each milestone.tasks as task}
			<li>
				<Task {task}></Task>
			</li>
		{/each}
	</ul>
</section>
