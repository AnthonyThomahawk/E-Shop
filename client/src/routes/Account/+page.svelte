<script lang="ts">
    import {onMount} from "svelte";
    import {getLocalStorage} from "../../lib/LocalStorage";
    import {goto} from '$app/navigation';

    let email = "";
    let notification = "";

    async function userLogOut() {
        localStorage.removeItem("UserData");
        notification = "User logged out, redirecting to home page...";
        await new Promise(resolve => setTimeout(resolve, 1000));
        await goto('/');
        location.reload();
    }

    onMount(async () => {
        try {
            const data = getLocalStorage("UserData");
            email = data.Email;
        }
        catch (error) {
            email = "";
        }
    });
</script>

<div>
    {#if email === ""}
        <h1>You are currently viewing as a guest</h1>
    {:else}
        <h1>You are currently logged in as:</h1>
        <h2>{email}</h2>
        <button on:click={userLogOut} style="width: 200px; height: 55px;">Log out</button>
        <h3>{notification}</h3>
    {/if}
</div>

<style>
    div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>