<script lang="ts">
    import {onMount} from "svelte";
    import {getLocalStorage, setLocalStorage} from "../lib/LocalStorage";
    import { page } from '$app/stores';

    let email = "";
    let currentPath = "";

    onMount(async () => {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);

        const oldheaderRefresh = getLocalStorage('refreshHeaderCount');

        if (oldheaderRefresh == undefined) {
            setLocalStorage('refreshHeaderCount', 0);
        } else {
            setLocalStorage('refreshHeaderCount', oldheaderRefresh + 1);
        }


        try {
            const data = getLocalStorage("UserData");

            email = data.Email;
        }
        catch (error) {
            email = "";
        }
    });

</script>

{#if email === ""}
    <h2>Welcome to world of tea, please log in or register to access all functions!</h2>
{:else}
    <h2>Welcome to world of tea, user {email} !</h2>
{/if}