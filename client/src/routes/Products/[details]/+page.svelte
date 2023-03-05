<script lang="ts">
    import {onMount} from "svelte";
    import {getProductDetails, getProducts} from "../../../lib/Products";
    import DynamicImage from "../../../lib/DynamicImage.svelte";
    import ShoppingCart from "../Images/shoppingCart.png"
    import PageData = App.PageData;

    export let data: PageData;

    let TargetProductID = data.post.id;

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

    let product: IProduct;
    onMount(async () => {
        const res = await getProductDetails(TargetProductID);

        product = res.map((item: any) => ({
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
    });

    let itemCount = 0;
</script>

<div class="flex-div">
    <h1>{product.Label}</h1>
    <h2>{product.Description}</h2>
    <div style="display: flex; flex-direction: row; align-items: start; position: relative;">
        <DynamicImage imageHeight=400 imageWidth=400 imageLink="{product.ImageURL}"/>
        <div style="padding-right: 50px;"></div>
        <div style="display: flex; flex-direction: column; position: relative;">
            <h3>{product.Characteristics}</h3>
            <h3><b>Current stock : {product.Stock}</b></h3>
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