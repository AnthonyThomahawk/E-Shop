<script lang="ts">
    import Header from '../lib/Header.svelte'
    import Footer from '../lib/Footer.svelte'
    import {getLocalStorage} from "../lib/LocalStorage";
    import {onMount} from "svelte";

    let refreshHeaderCount;
    let refreshHeader = 0;

    async function refresh() {
        refreshHeader += 1;
    }

    async function refresher() {
        while(true)
        {
            await new Promise(resolve => setTimeout(resolve, 500));
            refreshHeaderCount = getLocalStorage('refreshHeaderCount');
        }
    }

    onMount(async() => {
        await refresher();
    })

    $: {
        refreshHeaderCount;
        refresh();
    }
</script>

{#key refreshHeader}
    <Header/>
{/key}
<div style="padding:100px; align-items: center; justify-content: center">
    <slot/>
</div>
<Footer/>