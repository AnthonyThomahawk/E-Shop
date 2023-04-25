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

<div class="centered-div">
    <h1>Welcome to World Of Tea!</h1>
</div>
<div class="centered-div">
    <h2>This web application is a simple e-shop created from scratch<br>
        using Golang, PostgreSQL and Sveltekit :)
    </h2>
</div>
<div class="centered-div">
    <h2>You can start navigating the website by choosing either<br>
        the <a href="/Products" style="color:#0088ff">Products</a> page or the <a href="/Categories" style="color:#0088ff">Categories</a> page
    </h2>
</div>
<div class="centered-div">
    <h2>You may also use the top bar to navigate the website, <br>
    which allows you to change page at any time</h2>
</div>
<div class="centered-div">
    <h2>If you want to use the shopping cart functionality, you will have to create<br>
    a user account at the <a href="/Register" style="color:#0088ff">Register</a> page and then login to it using the <a href="/Login" style="color:#0088ff">Login</a> page.</h2>
</div>
<div class="centered-div">
    <h3>Powered by:</h3>
</div>
<div class="centered-div">
    <img src="static/assets/golang.png" height="100" width="200" alt="Golang">
    <div style="padding-left: 50px"></div>
    <img src="static/assets/postgresql.png" height="150" width="128" alt="PostgreSQL">
    <div style="padding-left: 50px"></div>
    <img src="static/assets/svelte.svg" height="150" width="128" alt="Sveltekit">
</div>

<style>
    .centered-div {
        display: flex;
        justify-content: center;
        text-align: center;
    }
</style>