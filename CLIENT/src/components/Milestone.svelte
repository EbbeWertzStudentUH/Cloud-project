<script>
	import Task from "./Task.svelte";

    let { milestone } = $props();
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
<section>
    <h2 class="mb-3 text-lg font-bold">Tasks: </h2>
				<ul class="space-y-2">
					{#each milestone.tasks as task}
						<li>
							<Task task={task}></Task>
						</li>
					{/each}
				</ul>
</section>