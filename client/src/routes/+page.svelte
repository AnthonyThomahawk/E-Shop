<script lang="ts">
    import {onMount} from "svelte";
    import {getLocalStorage} from "../lib/LocalStorage";

    let email = "";
    let token = "";

    onMount(async () => {
        try {
            const data = getLocalStorage("UserData");
            email = data.Email;
            token = data.Token;
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
