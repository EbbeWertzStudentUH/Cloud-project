<script>
	import '../app.css';


//   import { user, logout } from './+layout.js';
//   import { onMount } from 'svelte';

//   onMount(async () => {
//     await fetchUser();
//   });
// let user = {first_name:"dummy", last_name:"dummy"}
let user = null

function logout(){

}
</script>
<nav class="bg-slate-900 border-b-2 border-slate-500 text-white p-4 flex justify-between items-center">
  <div class="text-xl font-bold">ProjectCloud</div>
  <div class="text-l">Organise your projects on the cloud</div>
  <div>
    {#if user}
      <span class="mr-4">Hello, {user.first_name} {user.last_name}!</span>
      <button on:click={logout} class="bg-red-500 px-4 py-2 rounded">Logout</button>
    {:else}
      <button on:click={() => goto('/login')} class="bg-emerald-600 px-4 py-2 rounded">Login</button>
      <button on:click={() => goto('/register')} class="bg-emerald-800 px-4 py-2 rounded ml-2">Register</button>
    {/if}
  </div>
</nav>
<slot></slot>
