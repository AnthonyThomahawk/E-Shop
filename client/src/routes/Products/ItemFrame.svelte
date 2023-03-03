<script lang="ts">
    import DynamicImage from "../../lib/DynamicImage.svelte";
    export let ItemName = "Item name";
    export let ItemPrice = "Price";
    export let ItemID = 0;
    export let imageLink = "";
</script>

<div>
    <h2>{ItemName}</h2>
    <h3>Price: {ItemPrice}</h3>
    <DynamicImage imageLink="{imageLink}"/>
    <a href="/Products/{ItemID}">Buy now!</a>
</div>

<style>
    div {
        flex-direction: column;
        display:flex;
        justify-content:center;
        align-items: center;
        border: 2px solid black;
        border-radius:2%;
        width:300px;
        padding:10px;
        white-space: nowrap;
    }
</style>