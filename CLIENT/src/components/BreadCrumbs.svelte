<script>
    let { milestone, project } = $props()
</script>
<nav class="mb-4">
    <a href="/" class="text-emerald-500 hover:underline">Home</a>
    <span> / </span>
    {#if selectedMilestone}
        <button
            class="font-bold text-emerald-500"
            on:click={() => {
                selectedMilestone = null;
            }}>{current_open_project.name}</button
        >
        <span> / </span>
        <span class="font-bold text-slate-300">{selectedMilestone.name}</span>
    {:else}
        <span class="font-bold text-slate-300">{current_open_project.name}</span>
    {/if}
</nav>