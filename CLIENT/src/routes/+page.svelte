<script>
    import { page } from '$app/stores';
	import { addProject, projects_list } from '../stores/projects';
	import { writable } from 'svelte/store';
	import { fade, slide } from 'svelte/transition';
	import { onUpdateMessageType } from '../stores/updatemessages';
	import { user } from '../stores/user';

	let showForm = writable(false);
    let newProject = { name: '', deadline: '', github: '' };

	$:  current_projects_list = $projects_list;

	onUpdateMessageType("project_add", (subject, data) => {
		addProject(data);
	});

	function createProject(){
		showForm.set(false);
	}

</script>

<main class="flex justify-center items-center flex-grow h-full">
	<div class="bg-slate-800 border-2 border-emerald-700 p-8 shadow-md rounded w-[36rem]">
	  <h1 class="text-2xl font-bold mb-6 text-center text-emerald-500">Projects</h1>
	  	  <div class="mb-6 flex justify-center">
		<button
		  class="bg-emerald-600 hover:bg-emerald-700 text-white py-2 px-6 rounded shadow-md hover:shadow-emerald-400 transition-all duration-200"
		  on:click={() => showForm.set(true)}
		>
		  + Create New Project
		</button>
	  </div>
  
	  {#if $showForm}
	  <div in:slide={{ y: -20, duration:300}} out:slide={{ y: -20, duration:300}}>
		<div class="bg-slate-700 border border-emerald-600 p-4 rounded mb-6" in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
		  <h2 class="text-lg font-semibold mb-4 text-emerald-400">New Project</h2>
		  <form on:submit|preventDefault={createProject} class="space-y-4">
			<label class="block">
			  Name:
			  <input
				type="text"
				bind:value={newProject.name}
				class="p-2 w-full bg-slate-800 border border-slate-600 rounded text-white focus:outline-none focus:border-emerald-500"
				required
			  />
			</label>
			<label class="block">
			  Deadline:
			  <input
				type="date"
				bind:value={newProject.deadline}
				class="p-2 w-full bg-slate-800 border border-slate-600 rounded text-white focus:outline-none focus:border-emerald-500"
				required
			  />
			</label>
			<label class="block">
			  GitHub Repository:
			  <input
				type="url"
				bind:value={newProject.github}
				class="p-2 w-full bg-slate-800 border border-slate-600 rounded text-white focus:outline-none focus:border-emerald-500"
				required
			  />
			</label>
			<div class="flex justify-end space-x-4">
			  <button
				type="button"
				class="py-2 px-4 bg-gray-600 text-white rounded hover:bg-gray-700"
				on:click={() => showForm.set(false)}
			  >
				Cancel
			  </button>
			  <button
				type="submit"
				class="py-2 px-4 bg-emerald-600 text-white rounded hover:bg-emerald-700"
			  >
				Save
			  </button>
			</div>
		  </form>
		</div>
	</div>
  
		{:else}
		<div in:slide={{ y: -20, duration:300}} out:slide={{ y: -20, duration:300}}>
			<div in:fade={{ duration: 300 }} out:fade={{ duration: 200 }}>
		<ul class="space-y-4">
		  {#each $projects_list as project}
			<li>
			  <a 
				href={`/project/${project.id}`}
				class="block bg-slate-900 hover:bg-slate-800 text-white p-4 rounded shadow-lg hover:shadow-emerald-500 transition-all duration-200 transform hover:scale-105"
			  >
				<h2 class="text-lg font-semibold">{project.name}</h2>
			  </a>
			</li>
		  {/each}
		</ul>
	</div>
	</div>
		{/if}
	</div>

  </main>