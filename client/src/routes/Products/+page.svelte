<script lang="ts">
    import { onMount } from "svelte";
    import ItemFrame from "./ItemFrame.svelte";
    import { getProducts } from "../../lib/Products";
    import { each } from "svelte/internal";

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
        const res = await getProducts(1, 10);

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
    <ItemFrame
        ItemName="First item"
        ItemPrice="999$"
        ItemID="1"
        imagePath="./Images/svelte.svg"
    />
    <div style="padding:10px" />
    <ItemFrame
        ItemName="Second item"
        ItemPrice="50$"
        ItemID="2"
        imagePath="./Images/store-icon.svg"
    />
    <div style="padding:10px" />
    <ItemFrame
        ItemName="Third item"
        ItemPrice="15$"
        ItemID="3"
        imagePath="./Images/svelte.svg"
    />
</div>
