<script lang="ts">
    import {onMount} from "svelte";
    import {getProductDetails} from "../../../lib/Products";
    import {insertItem, updateItem, getCart} from "../../../lib/Cart";
    import DynamicImage from "../../../lib/DynamicImage.svelte";
    import ShoppingCart from "../Images/shoppingCart.png"
    // @ts-ignore
    import type {PageData} from './$types';
    import {getLocalStorage, setLocalStorage} from "../../../lib/LocalStorage";
    import {page} from "$app/stores";
    export let data: PageData;

    // @ts-ignore
    let targetID = data.post.id;

    interface IProductQty {
        ProductId: number;
        Quantity: number;
    }

    interface ICartItem {
        Product: IProduct;
        Quantity: number;
    }

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

    let product : IProduct = {} as IProduct;
    let imageSrc;

    let currentPath = "";

    let user;

    onMount(async () => {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);

        try {
            const data = getLocalStorage("UserData");
            user = data.Email;
        }
        catch (error) {
            user = "";
        }

        const [res, hd] = await getProductDetails(targetID);

        product = {
            ID : res.ID,
            CategoryID: res.CategoryID,
            SKU: res.SKU,
            Label: res.Label,
            Description: res.Description,
            Characteristics: res.Characteristics,
            Price: res.Price,
            Stock: res.Stock,
            ImageURL: res.ImageURL,
        };

        const img = new Image();
        img.src = product.ImageURL;
        imageSrc = img.src;
    });

    function truncate (num, places) {
        return Math.trunc(num * Math.pow(10, places)) / Math.pow(10, places);
    }

    let totalPrice = 0;
    let itemCount = 0;
    let cartLabel = "";

    async function addToCart() {
        let cartItems: Array<ICartItem>;

        const [res,hd] = await getCart(1,10);

        cartItems = res.Items.map((item: any) => ({
            Product : item.Product,
            Quantity : item.Quantity
        }));

        let itemFound = false;
        let oldQty = 0;

        for (let i = 0; i < cartItems.length; i++) {
            if (cartItems[i].Product.ID == targetID) {
                itemFound = true;
                oldQty = cartItems[i].Quantity;
                break;
            }
        }

        if (itemFound) {
            let newQty = oldQty + itemCount;
            await updateItem(targetID, newQty);
        } else {
            await insertItem(targetID, itemCount);
        }

        cartLabel = itemCount + " items added to cart!";
        itemCount = 0;
    }

    $: totalPrice = truncate(itemCount * product.Price, 1);
</script>

<div class="flex-div">
    <h1>{product.Label}</h1>
    <h2>{product.Description}</h2>
    <div style="display: flex; flex-direction: row; align-items: start; position: relative;">
        <img src="{imageSrc}" height=400px width=400px alt="Product Image"/>
        <div style="padding-right: 50px;"></div>
        <div style="display: flex; flex-direction: column; position: relative;">
            <h3>{product.Characteristics}</h3>
            <h3><b>Current stock : {product.Stock}</b></h3>
            <div style="display: flex;">
                <h3>Price : {product.Price}€</h3>
                <div style="padding-right: 15px"></div>
                {#if user !== ""}
                    <h3>Total price : {totalPrice}€</h3>
                {/if}
            </div>
            <div style="padding-bottom: 25px;"></div>

            {#if user !== ""}
                <div style="display: flex; flex-direction: row;">
                    <div style="display: flex; flex-direction: column;">
                        Quantity:
                        <input type=number bind:value={itemCount} min=0 max={product.Stock}>
                    </div>
                    <div style="padding-right: 30px;"></div>
                    <button on:click={addToCart} style="width:130px;height:50px;">
                        <div class="centered-div">
                            Add to cart
                            <div style="padding-right: 5px;"></div>
                            <DynamicImage imageHeight=25 imageWidth=25 imagePath="{ShoppingCart}"/>
                        </div>
                    </button>
                </div>
                <div style="padding-top: 30px;"></div>
                <div style="display: flex; flex-direction: row">
                    <h2>{cartLabel}</h2>
                    <div style="padding-left: 10px"></div>
                    {#if cartLabel !== ""}
                        <h2><a href="/Cart">Go to cart</a></h2>
                    {/if}
                </div>
            {:else}
                <h2>Log in to use the shopping cart!</h2>
            {/if}
        </div>
    </div>
</div>

<style>
    .flex-div {
        display: flex;
        flex-direction: column;
        justify-content: left;
        align-items: initial;
    }

    .centered-div {
        display:flex; justify-content: center; align-items: center
    }
</style>