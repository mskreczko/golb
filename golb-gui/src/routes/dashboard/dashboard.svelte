<script lang="js">
    const servers = [
        {addr: '192.168.0.1', healthcheckEndpoint: '/health', alive: true},
        {addr: '192.168.0.2', healthcheckEndpoint: '/health', alive: true},
        {addr: '192.168.0.3', healthcheckEndpoint: '/health', alive: false},
        {addr: '192.168.0.4', healthcheckEndpoint: '/health', alive: true},
        {addr: '192.168.0.5', healthcheckEndpoint: '/health', alive: false},
    ]
</script>

<style>
    .dot-red {
        height: 10px;
        width: 10px;
        background-color: #e32235;
        border: 1px solid #5e0c14;
        border-radius: 50%;
        display: inline-block;
    }

    .dot-green {
        height: 10px;
        width: 10px;
        background-color: #12fc21;
        border: 1px solid #1c7522;
        border-radius: 50%;
        display: inline-block;
    }

    .server-content {
        display: inline;
    }

    .server-block {
        border: 1px solid #000000;
        border-radius: 10px;
        padding: 5px;
        width: 200px;
        margin: 5px 0;
        box-shadow: 8px 8px 24px -11px rgba(66, 68, 90, 1);
    }

    .server-block:hover {
        background-color: #d9dbde;
        cursor: pointer;
    }
</style>

<h3>Servers</h3>
<div>
    {#each servers as server}
        <div class="server-block">
            <span class={server.alive ? 'dot-green' : 'dot-red'}></span>
            <p class="server-content">{server.addr}</p>
            <p class="server-content">{server.healthcheckEndpoint}</p>
        </div>
    {/each}
</div>