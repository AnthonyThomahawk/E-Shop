<script lang="ts">
    import Header from '../lib/Header.svelte'
    import Footer from '../lib/Footer.svelte'
    import {getLocalStorage} from "../lib/LocalStorage";
    import {onMount} from "svelte";
    import {getCart} from "../lib/Cart";
    import {goto} from "$app/navigation";

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

            if (localStorage.getItem("UserData") != null){
                try {
                    await getCart(1,10); // check if user is still authorized
                } catch (error) {
                    localStorage.removeItem("UserData");
                    await goto('/Login');
                    location.reload();
                }
            }
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