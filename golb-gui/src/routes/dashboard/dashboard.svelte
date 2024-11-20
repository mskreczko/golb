<script lang="js">
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
        background-color: #d9dbde;
        cursor: pointer;
    }
</style>

<ul class="list-none">
    {#each servers as server}
        <li class="server-block flex space-x-4 items-center w-full px-4 py-2 border-b border-gray-200 rounded-t-lg dark:border-gray-600">
            <span class={server.alive ? 'flex w-3 h-3 me-3 bg-green-500 rounded-full' : 'flex w-3 h-3 me-3 bg-red-500 rounded-full'}></span>
            <p>{server.addr}</p>
            <p>{server.healthcheckEndpoint}</p>
        </li>
    {/each}
</ul>