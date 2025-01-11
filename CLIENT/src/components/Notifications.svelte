<script>
    import { notifications, removeNotification } from "../stores/notifications";

    $: currentNotifications = $notifications;
</script>

<div class="fixed top-4 left-1/2 transform -translate-x-1/2 flex flex-col items-center space-y-4 z-50">
    {#each currentNotifications as notification}
        <button onclick={() => removeNotification(notification)} class="bg-emerald-800 rounded border-2 border-emerald-400 text-white w-3/4 max-w-md text-center">
            <h1 class="text-lg font-bold">{notification.message}</h1>
            <p class="text-base text-right pr-2">{notification.time}</p>
        </button>
    {/each}
</div>

