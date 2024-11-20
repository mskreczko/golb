<script lang="js">
    // const servers = [
    //     {addr: '192.168.0.1', healthcheckEndpoint: '/health', alive: true},
    //     {addr: '192.168.0.2', healthcheckEndpoint: '/health', alive: true},
    //     {addr: '192.168.0.3', healthcheckEndpoint: '/health', alive: false},
    //     {addr: '192.168.0.4', healthcheckEndpoint: '/health', alive: true},
    //     {addr: '192.168.0.5', healthcheckEndpoint: '/health', alive: false},
    // ]
    import {onMount} from "svelte";

    const endpoint = "http://127.0.0.1:8080/metrics/servers";
    let servers = [];

    onMount(async function () {
        const response = await fetch(endpoint);
        servers = await response.json();
    });
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