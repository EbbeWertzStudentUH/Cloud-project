<script>
	import { onMount } from 'svelte';
	import { addMilestoneToOpenProject, addTaskToMilestone, addUserToOpenProject, open_project } from '../../../stores/projects';
	import { friends } from '../../../stores/friends';
	import Project from '../../../components/Project.svelte';
	import Milestone from '../../../components/Milestone.svelte';
	import { slide } from 'svelte/transition';
	import { onUpdateMessageType } from '../../../stores/updatemessages';

	export let data;
	let { fetchProject, getGithubStats } = data;

	let selectedMilestone = null;
	let githubStats = null;

	$: current_open_project = $open_project;

	onMount(async () => {
		await fetchProject();
		if (current_open_project.github_repo != '') {
			githubStats = await getGithubStats(current_open_project.github_repo);
		}
	});
	onUpdateMessageType('milestone_add', (subject, data) => {
		addMilestoneToOpenProject(data);
	});
	onUpdateMessageType('user_add', (subject, data) => {
		addUserToOpenProject(data);
	});
	onUpdateMessageType('new_task_in_milestne', (subject, data) => {
		addTaskToMilestone(subject, data)
	});
</script>

<main class="p-4">
	{#if current_open_project}
		<nav class="mb-4">
			<a href="/" class="text-emerald-500 hover:underline">Home</a>
			<span> / </span>
			{#if selectedMilestone}
				<button
					class="font-bold text-emerald-500"
					onclick={() => {
						selectedMilestone = null;
					}}>{current_open_project.name}</button
				>
				<span> / </span>
				<span class="font-bold text-slate-300">{selectedMilestone.name}</span>
			{:else}
				<span class="font-bold text-slate-300">{current_open_project.name}</span>
			{/if}
		</nav>

		<div class="flex gap-4">
			<aside class="h-min w-1/4 rounded-2xl bg-slate-800 p-4 text-slate-200 shadow-lg">
				<h2 class="mb-3 text-lg font-bold">Milestones</h2>
				<ul class="space-y-2">
					{#each current_open_project.milestones as milestone}
						<li in:slide={{ y: -20, duration: 300 }}>
							<button
								class="w-full rounded-lg {selectedMilestone && milestone.id === selectedMilestone.id
									? 'translate-x-2 cursor-pointer border-2 border-emerald-400 bg-emerald-800'
									: 'bg-slate-700 hover:bg-emerald-500'} p-3 shadow transition-colors hover:text-slate-900"
								onclick={() => {
									selectedMilestone = milestone;
								}}
							>
								{milestone.name}
								<div class="gap-10 text-right">
									<table>
										<tbody>
											<tr>
												<td class="w-full pr-3 text-xs text-slate-500">Finished tasks: </td>
												<td class="font-bold text-emerald-200"
													>{milestone.num_of_finished_tasks}/{milestone.num_of_tasks}</td
												>
											</tr>
											<tr>
												<td class="w-full pr-3 text-xs text-slate-500">Problems: </td>
												<td class="font-bold text-red-200">{milestone.num_of_problems}</td>
											</tr>
										</tbody>
									</table>
								</div>
							</button>
						</li>
					{/each}
				</ul>
			</aside>

			<section class="flex-1 rounded-2xl bg-slate-900 p-6 text-slate-200 shadow-lg">
				{#if selectedMilestone}
					<div in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
						<Milestone milestone_id={selectedMilestone.id}></Milestone>
					</div>
				{:else}
					<div in:slide={{ y: -20, duration: 300 }} out:slide={{ y: -20, duration: 300 }}>
						<Project
							project={current_open_project}
							{githubStats}
							hasGithub={current_open_project.github_repo != ''}
						></Project>
					</div>
				{/if}
			</section>
		</div>
	{/if}
</main>
