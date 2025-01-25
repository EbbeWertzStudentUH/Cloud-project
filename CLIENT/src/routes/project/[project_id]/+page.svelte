<script>
	import { onMount } from 'svelte';
	import { open_project } from '../../../stores/projects';
    export let data;
    let {fetchProject} = data;

    $: current_open_project = $open_project;
    onMount(async () => {
		await fetchProject();
    });
</script>

<main class="p-4">
    <nav class="mb-4">
        <a href="/" class="text-emerald-500 hover:underline">Home</a>
        <span> / </span>
        {#if current_open_project}
            <span class="font-bold">{current_open_project.name}</span>
        {/if}
    </nav>
</main>
