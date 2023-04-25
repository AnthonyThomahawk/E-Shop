<script lang="ts">
    import {onMount} from "svelte";

    export let currentPage = 1;
    export let maxPages = 5;

    let pageList = [];

    for (let i = 1; i <= maxPages; i++) {
        pageList.push(i);
    }

    async function updatePage(pn) {
        const pageNumber = +pn;

        if (currentPage == 1 && pageNumber == 0) {
            return;
        }
        if (currentPage == maxPages && pageNumber == maxPages+1) {
            return;
        }

        currentPage = pageNumber;
    }
</script>

<button on:click={() => updatePage(`${currentPage-1}`)} style="width: 3vw; height:2vw">←</button>
<div style="padding-left: 7px;"></div>
{#each pageList as p}
    {#if p === currentPage}
        <button on:click={() => updatePage(`${p}`)} style="width: 2vw; height:2vw; color:red; border-radius: 100%">{p}</button>
    {:else}
        <button on:click={() => updatePage(`${p}`)} style="width: 2vw; height:2vw; color:black; border-radius: 100%">{p}</button>
    {/if}
    <div style="padding-left: 7px;"></div>
{/each}
<button on:click={() => updatePage(`${currentPage+1}`)} style="width: 3vw; height:2vw">→</button>