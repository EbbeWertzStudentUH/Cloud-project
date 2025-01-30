<script>
	import { user } from '../stores/user';
	import { open_project } from '../stores/projects';
	import { assignToTask, completeTask, makeProblemInTask, ResolveProblemInTask } from '$lib/project_requests';
	import { fade, slide, fly } from 'svelte/transition';

	let { task } = $props();
	let currUser = $user;
	let showForm = $state(false);
	let newProblem = $state({ name: '' });

	async function assignTask() {
		await assignToTask($open_project.id, task.id);
	}
	async function setTaskComplete() {
		await completeTask($open_project.id, task.id);
	}
	async function createProblem() {
		showForm = false;
        await makeProblemInTask($open_project.id, task.id, newProblem.name);
	}
    async function resolveProblem(id) {
		showForm = false;
        await ResolveProblemInTask($open_project.id, task.id, id);
	}
</script>

<section class="mb-6 flex justify-between rounded-lg bg-slate-800 p-4 shadow-md">
	<header>
		<h3 class="mb-4 text-xl font-bold text-emerald-600">{task.name}</h3>
		{#if task.status == 'open'}
			<span class="text-md rounded-md bg-slate-700 px-4 py-2 font-bold">Not assigned yet</span>
			<div class="m-4"></div>
			<button
				class="m-2 rounded-md bg-emerald-600 px-8 text-lg text-slate-900 hover:bg-emerald-400 hover:text-emerald-900"
				onclick={assignTask}>Assign yourself to Task</button
			>
		{:else if task.status == 'active'}
			<span class="text-md rounded-md bg-emerald-900 px-4 py-2 font-bold">Active</span>
			<ul class="m-4 space-y-2">
				<li>
					<span class="text-sm text-slate-500">Assigned to: </span>
					<span class="font-bold">{task.user.first_name}</span>
				</li>

				<li>
					<span class="text-sm text-slate-500">Since: </span>
					<span class="font-bold">{task.active_period_start}</span>
				</li>
			</ul>
			{#if currUser.id == task.user.id}
				<button
					class="m-2 rounded-md bg-emerald-600 px-8 text-lg text-slate-900 hover:bg-emerald-400 hover:text-emerald-900"
					onclick={setTaskComplete}>Finish Task</button
				>
			{/if}
		{:else if task.status == 'closed'}
			<span class="text-md rounded-md bg-slate-600 px-4 py-2 font-bold">Finished</span>
			<ul class="m-4 space-y-2">
				<li>
					<span class="text-sm text-slate-500">Assigned to: </span>
					<span class="font-bold">{task.user.first_name}</span>
				</li>
				<li>
					<span class="text-sm text-slate-500">From: </span>
					<span class="font-bold">{task.active_period_start}</span>
				</li>
				<li>
					<span class="text-sm text-slate-500">Untill: </span>
					<span class="font-bold">{task.active_period_end}</span>
				</li>
			</ul>
		{/if}
	</header>

    {#if task.status == 'active' }
        
    
	<section class="mb-6 rounded-lg bg-slate-800 p-4 min-w-64">
        
        <h4 class="mb-3 text-xl font-semibold">Problems ({task.num_of_problems})</h4>
		{#if showForm}
			<section
				class="mb-6 rounded border border-emerald-600 bg-slate-700 p-4"
				in:slide={{ y: -20, duration: 300 }}
				out:slide={{ y: -20, duration: 300 }}
			>
				<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
					<h2 class="mb-4 text-lg font-semibold text-emerald-400">New Problem</h2>
					<form onsubmit={createProblem} class="space-y-4">
						<label class="block">
							Name:
							<input
								type="text"
								bind:value={newProblem.name}
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
			<section in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }} class="-translate-y-9 flex justify-end w-full bg-red-50 h-0">
				<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
                    
					<button
						onclick={() => {
							showForm = true;
						}}
						class="rounded-md bg-emerald-600 px-6 text-slate-900 shadow transition-colors hover:bg-emerald-600"
					>
						Add
					</button>
				</div>
			</section>
		{/if}

		<ul class="space-y-2">
			{#each task.problems as problem}
				<li class="border-b-2 border-slate-900">
					<span class="pr-16 font-bold text-red-900">{problem.name}</span>
					<div class="flex justify-between">
						<button
							class="m-2 rounded-md bg-emerald-600 px-8 text-slate-900 hover:bg-emerald-400 hover:text-emerald-900"
							onclick={async () =>  { await resolveProblem(problem.id)}}
                            >resolve</button
						>
						<span class="text-right text-sm text-slate-500">{problem.posted_at}</span>
					</div>
				</li>
			{/each}
		</ul>
	</section>
    {/if}
</section>
