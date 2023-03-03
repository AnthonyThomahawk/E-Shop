<script lang="ts">
    import {onMount} from "svelte";
    import {getProducts} from "../../../lib/Products";
    import DynamicImage from "../../../lib/DynamicImage.svelte";

    /** @type {import('./$types').PageData} */
    export let data;


    interface IProduct {
        CategoryID: number;
        SKU: string;
        Label: string;
        Description: string;
        Characteristics: string;
        Price: number;
        Stock: number;
        ImageURL: string;
    }

    let products: Array<IProduct> = [];
    onMount(async () => {
        const res = await getProducts(1, 5);

        products = res.map((item: any) => ({
            CategoryID: item.CategoryID,
            SKU: item.SKU,
            Label: item.Label,
            Description: item.Description,
            Characteristics: item.Characteristics,
            Price: item.Price,
            Stock: item.Stock,
            ImageURL: item.ImageURL,
        }));
    });
</script>

<!--temporary code until getProductFromSKU() is implemented-->
{#each products as product}
    {#if product.SKU === data.post.id}
        <div class="flex-div">
            <h1>{product.Label}</h1>
            <h2>{product.Description}</h2>
            <h3>{product.Characteristics}</h3>
        </div>
        <div class="centered-div">
            <DynamicImage imageHeight=400 imageWidth=400 imageLink="{product.ImageURL}"/>
        </div>

    {/if}
{/each}

<style>
    .flex-div {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .centered-div {
        display:flex; justify-content: center; align-items: center
    }
</style>