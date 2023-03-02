<script lang="ts">
    import { onMount } from "svelte";

    export let ItemName = "Item name";
    export let ItemPrice = "Price";
    export let ItemID = 0;
    export let imagePath = ""; // pass either a path to a local image, or a URL in the WWW
    export let imageLink = "";
    let imageSrc = "";

    async function loadImage() {
        if (imagePath !== "")
        {
            const module = await import(imagePath);
            imageSrc = module.default;
        }
        else if (imageLink !== "")
        {
            const img = new Image();
            img.src = imageLink;
            imageSrc = img.src;
        }
    }

    onMount(loadImage);
</script>

<div>
    <h2>{ItemName}</h2>
    <h3>Price: {ItemPrice}</h3>
    <img src={imageSrc} alt="" style="width:128px; height:128px; padding-bottom: 10px"/>
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