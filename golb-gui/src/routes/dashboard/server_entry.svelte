<script>
    let {addr, alive} = $props();
    let expanded = $state(false);
</script>

<div class="collapsible w-full">
    <h3 class="flex items-center">
        <button class="flex items-center w-full" aria-expanded={expanded} onclick={() => expanded = !expanded}>
            <div class="flex items-center">
                <span class={alive ? 'flex w-3 h-3 me-3 bg-green-500 rounded-full' : 'flex w-3 h-3 me-3 bg-red-500 rounded-full'}></span>
                {addr}
            </div>
            <div>
                <svg viewBox="0 0 20 20" fill="none">
                    <path class="vert" d="M10 1V19" stroke="black" stroke-width="2"/>
                    <path d="M1 10L19 10" stroke="black" stroke-width="2"/>
                </svg>
            </div>
        </button>
    </h3>

    <div class='contents' hidden={!expanded}>
        <slot></slot>
    </div>
</div>

<style>
    .collapsible {
        border-bottom: 1px solid var(--gray-light, #eee);
    }

    h3 {
        margin: 0;
    }

    button {
        background-color: var(--background, #fff);
        color: var(--gray-darkest, #282828);
        display: flex;
        justify-content: space-between;
        width: 100%;
        border: none;
        margin: 0;
        padding: 1em 0.5em;
    }

    button[aria-expanded="true"] {
        border-bottom: 1px solid var(--gray-light, #eee);
    }

    button[aria-expanded="true"] .vert {
        display: none;
    }

    button:focus svg {
        outline: 2px solid;
    }

    svg {
        height: 0.7em;
        width: 0.7em;
    }
</style>