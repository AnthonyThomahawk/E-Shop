<script>
    import {onMount} from "svelte";
    import Icon from '../assets/TeaIcon.png'
    import {getLocalStorage} from "../lib/LocalStorage";

    export let title = "World of tea";
    let user;

    onMount(async () => {
        try {
            const data = getLocalStorage("UserData");
            user = data.Email;
        }
        catch (error) {
            user = "";
        }
    });
</script>

<header>
    <img src={Icon} alt="Main icon">
    <h1>{title}</h1>
    <nav style="padding-left:20px;">
        <a href="/" style="color:#fff">home</a>
        <a href="/Categories" style="color:#fff">categories</a>
        <a href="/Products" style="color:#fff">products</a>
    </nav>
    {#if user === ""}
        <div style="padding-right: 55px; margin-left: auto;">
            <a href="/Login" style="color:#fff">login</a>
            <a href="/Register" style="padding-left:10px; color:#fff">register</a>
        </div>
    {:else}
        <div style="padding-right: 55px; margin-left: auto;">
            <a href="/Cart" style="color:#fff">cart</a>
            <a href="/Account" style="padding-left:10px; color:#fff">my account</a>
        </div>
    {/if}
</header>

<style>
    img {
        height:50px;
        width:50px;
        padding-right:10px;
    }
    header {
        font-size: 165%;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;
        background-color: #333;
        color: #fff;
        height: 100px;
        padding: 0 16px;
    }

    h1 {
        margin: 0;
    }
</style>
