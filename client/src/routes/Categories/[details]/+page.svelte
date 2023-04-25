<script lang="ts">
    import {getProductsByCategoryId} from "../../../lib/Products";
    import ItemFrame from "../../Products/ItemFrame.svelte";
    // @ts-ignore
    import type {PageData} from './$types';
    import {page} from "$app/stores";
    import {setLocalStorage} from "../../../lib/LocalStorage";
    export let data: PageData;

    // @ts-ignore
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

    let pageNumber = 1;
    let pageChanges = 0;

    let currentPath;

    let products: Array<IProduct> = [];

    async function getProductList()
    {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);

        const [res] = await getProductsByCategoryId(pageNumber, 5, CategoryID);

        products = res.map((item: any) => ({
            ID : item.ID,
            CategoryID: item.CategoryID,
            SKU: item.SKU,
            Label: item.Label,
            Description: item.Description,
            Characteristics: item.Characteristics,
            Price: item.Price,
            Stock: item.Stock,
            ImageURL: item.ImageURL,
        }));

        pageChanges += 1;
    }


    $: {
        pageNumber;
        getProductList();
    }
</script>

<h1 style="display: flex; justify-content: center">Products</h1>

<div style="display:flex;">
    {#key pageChanges}
        {#each products as product}
            <ItemFrame
                    ItemName="{product.Label}"
                    ItemPrice="{product.Price}€"
                    ItemID="{product.ID}"
                    imageLink="{product.ImageURL}"
            />
            <div style="padding:10px"></div>
        {/each}
    {/key}
</div>