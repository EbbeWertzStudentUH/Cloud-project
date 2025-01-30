<script>
	import { user } from "../stores/user";
	import { open_project } from '../stores/projects';
	import { assignToTask, completeTask } from "$lib/project_requests";

    let { task } = $props();
    let currUser = $user;
    async function assignTask() {
		await assignToTask($open_project.id, task.id);
	}
    async function setTaskComplete() {
		await completeTask($open_project.id, task.id);
	}
</script>
<section class="mb-6 rounded-lg bg-slate-800 p-4 shadow-md flex justify-between">
    <header>
        <h3 class="mb-4 text-xl font-bold text-emerald-600">{task.name}</h3>
        {#if task.status == "open" }
        <span class="bg-slate-700 font-bold text-md rounded-md py-2 px-4">Not assigned yet</span>
        <div class="m-4"></div>
        <button class="bg-emerald-600 rounded-md m-2 px-8 text-lg hover:bg-emerald-400 hover:text-emerald-900" onclick={assignTask}>Assign yourself to Task</button>
        {:else if task.status == "active" }
        <span class="bg-emerald-900 font-bold text-md rounded-md py-2 px-4">Active</span>
        <ul class="space-y-2 m-4">
            <li>
                <span class="text-slate-500 text-sm">Assigned to: </span>
                <span class="font-bold">{task.user.first_name}</span>
            </li>
            
            <li>
                <span class="text-slate-500 text-sm">Since: </span>
                <span class="font-bold">{task.active_period_start}</span>
            </li>
        </ul>
        {#if currUser.id == task.user.id }
            <button class="bg-emerald-600 rounded-md m-2 px-8 text-lg hover:bg-emerald-400 hover:text-emerald-900" onclick={setTaskComplete} >Finish Task</button>
        {/if}
        {:else if task.status == "closed" }
        <span class="bg-slate-600 font-bold text-md rounded-md py-2 px-4">Finished</span>
        <ul class="space-y-2 m-4">
            <li>
                <span class="text-slate-500 text-sm">Assigned to: </span>
                <span class="font-bold">{task.user.first_name}</span>
            </li>
            <li>
                <span class="text-slate-500 text-sm">From: </span>
                <span class="font-bold">{task.active_period_start}</span>
            </li>
            <li>
                <span class="text-slate-500 text-sm">Untill: </span>
                <span class="font-bold">{task.active_period_end}</span>
            </li>
        </ul>
        {/if}
        </header>
    <section class="mb-6 rounded-lg bg-slate-800 p-2">
        <h4 class="mb-3 text-xl font-semibold">Problems ({task.num_of_problems})</h4>
        <ul class="space-y-2">
            {#each task.problems as problem }
            <li class="border-b-2 border-slate-900 ">
                <span class="text-red-900 font-bold">{problem.name}</span>
                <div class="flex justify-between">
                    <button class="bg-emerald-600 rounded-md m-2 px-8 hover:bg-emerald-400 hover:text-emerald-900">resolve</button>
                    <span class="text-right text-sm text-slate-500">{problem.posted_at}</span>
                </div>
            </li>
            {/each}
        </ul>
    </section>
</section>
