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
        const data = getLocalStorage("UserData");
        email = data.Email;
    });
</script>

<h1>You are currently logged in as:</h1>
<h2>{email}</h2>
<button on:click={userLogOut} style="width: 200px; height: 55px;">Log out</button>
<h3>{notification}</h3>
