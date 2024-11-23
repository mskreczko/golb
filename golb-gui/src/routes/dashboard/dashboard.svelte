<script lang="js">
    import ServerEntry from './server_entry.svelte';
    import ServerDetails from './server_details.svelte';
    const endpoint = "http://127.0.0.1:8080/metrics/servers";
    let servers = [];

    async function getServers() {
        const response = await fetch(endpoint);
        servers = await response.json();
    }

    let clear
    $: {
        clearInterval(clear)
        clear = setInterval(getServers, 5000)
    }
</script>

<style>
    .server-block:hover {
        cursor: pointer;
    }
</style>

<ul class="list-none">
    {#each servers as server}
        <li class="server-block flex space-x-4 items-center w-full px-4 py-2 border-b border-gray-200 rounded-t-lg dark:border-gray-600">
            <ServerEntry addr={server.addr} alive={server.alive}>
                <ServerDetails healthcheckEndpoint={server.healthcheckEndpoint} lastAlive={server.lastAlive} />
            </ServerEntry>
        </li>
    {/each}
</ul>