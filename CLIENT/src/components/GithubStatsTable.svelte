<script>
    let { githubStats } = $props();
    function getNumOrEmoji(num){
        if(num == 1) return '🥇';
        if(num == 2) return '🥈';
        if(num == 3) return '🥉';
        return num;
    }
</script>
{#if githubStats && githubStats.length > 0}
  <section class="mb-6 rounded-lg bg-slate-800 p-4 shadow-md">
    <h2 class="mb-4 text-xl font-semibold text-emerald-400">GitHub Stats</h2>
    <div class="overflow-x-auto">
      <table class="w-full border-collapse text-left text-sm text-slate-400">
        <thead>
            <tr class="bg-slate-900 text-emerald-400">
                <th rowspan="2" class="p-3">#</th>
                <th rowspan="2" class="p-3">Username</th>
                <th colspan="3" class="p-3 text-center">Total</th>
                <th colspan="3" class="p-3 text-center">This Week</th>
              </tr>
              <!-- Second level of headings -->
              <tr class="bg-slate-900 text-emerald-400">
                <th class="p-3"><span class="p-2 bg-blue-950 rounded-lg">Commits</span></th>
                <th class="p-3"><span class="p-2 bg-green-950 rounded-lg">Additions</span></th>
                <th class="p-3"><span class="p-2 bg-red-950 rounded-lg">Deletions</span></th>
                <th class="p-3"><span class="p-2 bg-blue-950 rounded-lg">Commits</span></th>
                <th class="p-3"><span class="p-2 bg-green-950 rounded-lg">Additions</span></th>
                <th class="p-3"><span class="p-2 bg-red-950 rounded-lg">Deletions</span></th>
              </tr>
        </thead>
        <tbody>
          {#each githubStats as stat, index}
            <tr class="border-b border-slate-700 hover:bg-slate-700">
              <td class="p-3">{getNumOrEmoji(index + 1)}</td>
              <td class="p-3">{stat.username}</td>
              <td class="p-3">{stat.totalCommits}</td>
              <td class="p-3">{stat.totalAdditions}</td>
              <td class="p-3">{stat.totalDeletions}</td>
              <td class="p-3">{stat.lastWeekCommits}</td>
              <td class="p-3">{stat.lastWeekAdditions}</td>
              <td class="p-3">{stat.lastWeekDeletions}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </section>
{:else}
  <section class="mb-6 rounded-lg bg-slate-800 p-4 text-center shadow-md">
    <p class="text-slate-400">No GitHub stats available.</p>
  </section>
{/if}
