<script lang="ts">
    import { onMount } from "svelte";
    import ItemFrame from "./ItemFrame.svelte";
    import { getProducts } from "../../lib/Products";

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

<h1>Products</h1>
<div style="display:flex;">
    {#each products as product}
        <ItemFrame
                ItemName="{product.Label}"
                ItemPrice="{product.Price}€"
                ItemID="{product.SKU}"
                imageLink="{product.ImageURL}"
        />
        <div style="padding:10px" />
    {/each}
</div>
