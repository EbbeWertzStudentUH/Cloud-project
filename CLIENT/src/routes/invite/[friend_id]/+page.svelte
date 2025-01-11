<script>
	import { goto } from '$app/navigation';

    export let data;
    const { friend } = data;

    async function sendFriendRequest() {
        await data.sendRequest(friend.id);
        goto('/');
    }
</script>

<main class="flex justify-center items-center flex-grow h-full">
  <div class="bg-slate-800 border-2 border-emerald-700 p-8 shadow-md rounded w-96">
    <h1 class="text-xl font-bold mb-4">Friend Request</h1>
    <p class="text-white mb-6">
      Do you want to send a friend request to <span class="font-semibold">{friend.first_name} {friend.last_name}</span>?
    </p>
    <button 
      on:click={sendFriendRequest}
      class="w-full bg-emerald-600 hover:bg-emerald-700 text-white py-2 rounded transition-colors duration-200">
      Send Friend Request
    </button>
  </div>
</main>
