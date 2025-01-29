<script>
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
    <pre>{JSON.stringify(milestone, null, 2)}</pre>
</header>