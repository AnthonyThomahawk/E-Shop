<script lang="ts">
    import {onMount} from "svelte";
    import {getProductDetails} from "../../../lib/Products";
    import DynamicImage from "../../../lib/DynamicImage.svelte";
    import ShoppingCart from "../Images/shoppingCart.png"
    // @ts-ignore
    import type {PageData} from './$types';
    export let data: PageData;

    // @ts-ignore
    let TargetID = data.post.id;

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

    onMount(async () => {
        const res = await getProductDetails(TargetID);

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

    $ : totalPrice = truncate(itemCount * product.Price, 1);
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
                <h3>Total price : {totalPrice}€</h3>
            </div>
            <div style="padding-bottom: 25px;"></div>
            <div style="display: flex; flex-direction: row;">
                <div style="display: flex; flex-direction: column;">
                    Quantity:
                    <input type=number bind:value={itemCount} min=0 max={product.Stock}>
                </div>
                <div style="padding-right: 30px;"></div>
                <button style="width:130px;height:50px;">
                    <div class="centered-div">
                        Add to cart
                        <div style="padding-right: 5px;"></div>
                        <DynamicImage imageHeight=25 imageWidth=25 imagePath="{ShoppingCart}"/>
                    </div>
                </button>
            </div>
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