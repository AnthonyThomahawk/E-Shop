<script lang="ts">
    import ItemFrame from "./ItemFrame.svelte";
    import {getProducts} from "../../lib/Products";
    import {page} from "$app/stores";
    import {setLocalStorage} from "../../lib/LocalStorage";
    import Pagination from "../../lib/Pagination.svelte";
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

    let maxPages = 5;
    let pageNumber = 1;
    let pageChanges = 0;
    let currentPath = "";

    let products: Array<IProduct> = [];

    async function getProductList()
    {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);
        const [res] = await getProducts(pageNumber, 5);

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
<div style="display: flex; flex-direction: row; align-items: center; justify-content: center">
    <div style="padding:5px"></div>
    <Pagination bind:currentPage={pageNumber} maxPages={maxPages}></Pagination>
</div>
<div style="padding:10px"></div>
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
