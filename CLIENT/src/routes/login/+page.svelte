<script>
  import { goto } from '$app/navigation';
	import { POST } from '$lib';

  export let data;

  let email = '';
  let password = '';
  let error = '';

  async function login() {
    const returneduser = await POST({email, password}, 'user/authenticate');
    if(returneduser.valid){
      data.user.set(returneduser);
      goto('/');
    } else {
      error = 'invalid credidentials'
    }
  }
</script>

<main class="flex justify-center items-center flex-grow">
  <div class="bg-slate-800 border-2 border-emerald-700 p-8 shadow-md rounded w-96">
    <h1 class="text-xl font-bold mb-4">Login</h1>
    {#if error}
      <p class="text-red-500 mb-4">{error}</p>
    {/if}
    <form on:submit|preventDefault={login}>
      <label class="block mb-2">
        Email:
        <input type="email" class="{error? "border-4  border-red-500" : ""} p-2 w-full rounded text-slate-950" bind:value={email} required />
      </label>
      <label class="block mb-4">
        Password:
        <input type="password" class="{error? "border-4  border-red-500" : ""} p-2 w-full rounded text-slate-950" bind:value={password} required />
      </label>
      <button type="submit" class="w-full bg-emerald-600 text-white py-2 rounded">Login</button>
    </form>
  </div>
</main>
