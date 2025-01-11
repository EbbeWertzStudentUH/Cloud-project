<script>
    import { goto } from '$app/navigation';
    import { POST } from '$lib';
    import { updateUser } from '../../stores/user.js';

  
    export let data;
  
    let email = '';
    let password = '';
    let first_name = '';
    let last_name = '';
  
    async function register() {
      const resp = await POST({email, password, first_name, last_name}, '/user/create_account');
      if(resp.valid){
        localStorage.setItem('authToken', resp.token);
        const { id, first_name, last_name } = data.user;
			  updateUser({ id, first_name, last_name });
        goto('/');
      } else {
        error = 'invalid credidentials'
      }
    }
  </script>
  
  <main class="flex justify-center items-center flex-grow">
    <div class="bg-slate-800 border-2 border-emerald-700 p-8 shadow-md rounded w-96">
      <h1 class="text-xl font-bold mb-4">Register</h1>
      <form on:submit|preventDefault={register}>
        <label class="block mb-2">
            Fist name:
            <input type="text" class="p-2 w-full rounded text-slate-950" bind:value={first_name} required />
          </label>
          <label class="block mb-2">
            Last name:
            <input type="text" class="p-2 w-full rounded text-slate-950" bind:value={last_name} required />
          </label>
        <label class="block mb-2">
          Email:
          <input type="email" class="p-2 w-full rounded text-slate-950" bind:value={email} required />
        </label>
        <label class="block mb-4">
          Password:
          <input type="password" class="p-2 w-full rounded text-slate-950" bind:value={password} required />
        </label>
        <button type="submit" class="w-full bg-emerald-600 text-white py-2 rounded">Register</button>
      </form>
    </div>
  </main>
  