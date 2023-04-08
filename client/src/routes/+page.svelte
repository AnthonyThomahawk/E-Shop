<script lang="ts">
    import {onMount} from "svelte";
    import {getLocalStorage, setLocalStorage} from "../lib/LocalStorage";
    import {goto} from "$app/navigation";
    import { page } from '$app/stores';

    let email = "";
    let currentPath = "";

    onMount(async () => {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);
        setLocalStorage('refreshHeaderCount', 0);

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