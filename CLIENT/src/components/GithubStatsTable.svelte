<script>
	import { text } from '@sveltejs/kit';
	import { fade, slide } from 'svelte/transition';

	let { githubStats } = $props();

	function getNumOrEmoji(num) {
		if (num == 1) return '🥇';
		if (num == 2) return '🥈';
		if (num == 3) return '🥉';
		return '#' + num;
	}
</script>

{#if githubStats && githubStats.length > 0}
	<section in:fade={{ duration: 300 }}
	class="mb-6 rounded-lg bg-slate-800 p-4 shadow-md">
		<h2 class="mb-4 text-xl font-semibold text-emerald-400">GitHub Stats</h2>
		<div class="overflow-x-auto">
			<table class="w-full border-collapse text-left text-sm text-slate-400">
				<thead>
					<tr class="bg-slate-900 text-emerald-400">
						<th rowspan="2" class="p-3">#</th>
						<th rowspan="2" class="p-3">Username</th>
						<th colspan="3" class="p-3 text-center">Total</th>
						<th colspan="3" class="p-3 text-center">This Week</th>
					</tr>
					<!-- Second level of headings -->
					<tr class="bg-slate-900 text-emerald-400">
						<th class="p-3"><span class="rounded-lg bg-blue-950 p-2">Commits</span></th>
						<th class="p-3"><span class="rounded-lg bg-green-950 p-2">Additions</span></th>
						<th class="p-3"><span class="rounded-lg bg-red-950 p-2">Deletions</span></th>
						<th class="p-3"><span class="rounded-lg bg-blue-950 p-2">Commits</span></th>
						<th class="p-3"><span class="rounded-lg bg-green-950 p-2">Additions</span></th>
						<th class="p-3"><span class="rounded-lg bg-red-950 p-2">Deletions</span></th>
					</tr>
				</thead>
				<tbody>
					{#each githubStats as stat, index}
						<tr class="border-b border-slate-700 hover:bg-slate-700">
							<td class="pl-3 text-2xl">{getNumOrEmoji(index + 1)}</td>
							<td class="p-3">{stat.username}</td>
							<td class="p-3 font-extrabold text-emerald-500"
								>{new Intl.NumberFormat().format(stat.totalCommits)}</td
							>
							<td class="p-3">{new Intl.NumberFormat().format(stat.totalAdditions)}</td>
							<td class="p-3">{new Intl.NumberFormat().format(stat.totalDeletions)}</td>
							<td class="p-3">{new Intl.NumberFormat().format(stat.lastWeekCommits)}</td>
							<td class="p-3">{new Intl.NumberFormat().format(stat.lastWeekAdditions)}</td>
							<td class="p-3">{new Intl.NumberFormat().format(stat.lastWeekDeletions)}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</section>
{:else}
	<section out:slide={{ duration: 300 }} class="mb-6 animate-pulse rounded-lg bg-slate-800 p-4 text-center shadow-md">
		<p class="text-2xl font-bold text-slate-600">Github statistics</p>
		<div class="flex items-center justify-center space-x-2">
			{#each [0, 1, 2] as dot}
				<div
					class="my-4 h-2 w-2 animate-bounce rounded-full bg-slate-400 [animation-delay:-0.{dot *
						3}s]"
				></div>
			{/each}
		</div>
	</section>
{/if}
