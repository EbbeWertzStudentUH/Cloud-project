<script>
  import { goto } from '$app/navigation';

  let email = '';
  let password = '';
  let error = '';

  async function login() {
  }
</script>

<main class="flex justify-center items-center min-h-screen bg-gray-100">
  <div class="bg-white p-8 shadow-md rounded w-96">
    <h1 class="text-xl font-bold mb-4">Login</h1>
    {#if error}
      <p class="text-red-500 mb-4">{error}</p>
    {/if}
    <form on:submit|preventDefault={login}>
      <label class="block mb-2">
        Email:
        <input type="email" class="border border-gray-300 p-2 w-full" bind:value={email} required />
      </label>
      <label class="block mb-4">
        Password:
        <input type="password" class="border border-gray-300 p-2 w-full" bind:value={password} required />
      </label>
      <button type="submit" class="w-full bg-blue-500 text-white py-2 rounded">Login</button>
    </form>
  </div>
</main>
