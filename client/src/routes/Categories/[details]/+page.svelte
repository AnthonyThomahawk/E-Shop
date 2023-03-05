<script lang="ts">
    import {onMount} from "svelte";
    import {getProductsByCategoryId} from "../../../lib/Products";
    import ItemFrame from "../../Products/ItemFrame.svelte";
    import PageData = App.PageData;
    export let data: PageData;

    let CategoryID = data.post.id;

    interface IProduct {
        ID: number;
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
        const res = await getProductsByCategoryId(1, 5, CategoryID);

        products = res.map((item: any) => ({
            ID: item.ID,
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
                ItemID="{product.ID}"
                imageLink="{product.ImageURL}"
        />
        <div style="padding:10px" />
    {/each}
</div>