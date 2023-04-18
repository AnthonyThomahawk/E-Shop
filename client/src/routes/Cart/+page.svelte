<script lang="ts">
    import {onMount} from "svelte";
    import {getLocalStorage} from "../../lib/LocalStorage";
    import {getCart, updateItem, deleteItem} from "../../lib/Cart";

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

    interface ICartItem {
        Product: IProduct;
        Quantity: number;
    }

    let user;
    let cartItems: Array<ICartItem> = [];
    let total: number;
    let cartChanges = 0;
    let cartChangesCompleted = 0;

    async function getCartItems() {
        try {
            const [res] = await getCart(1,10);

            cartItems = res.Items.map((item: any) => ({
                Product : item.Product,
                Quantity : item.Quantity
            }));

            cartChangesCompleted++;
        }
        catch (error) {
            console.log(error);
        }
    }

    onMount(async () => {
        try {
            const data = getLocalStorage("UserData");
            user = data.Email;
        }
        catch (error) {
            user = "";
        }

        await getCartItems();
    });

    $: {
        cartItems;

        total = 0;

        for (let i = 0; i < cartItems.length; i++) {
            if (cartItems[i].Quantity < 1)
                continue;
            total += cartItems[i].Quantity * cartItems[i].Product.Price;
            updateItem(cartItems[i].Product.ID, cartItems[i].Quantity);
        }
    }

    $: {
        cartChanges;
        getCartItems();
    }

    async function del(productID: string) {
        await deleteItem(+productID);
        cartChanges++;
    }

    function roundNumber(num, decimals) {
        return (Math.round(num * 100) / 100).toFixed(decimals);
    }
</script>

{#if user !== ""}
    {#if total !== 0}
        <div style="align-items: center">
            <h1>Total: {roundNumber(total, 2)} €</h1>
        </div>
        <div class="container">
            <table>
                <thead>
                <tr>
                    <th>Product</th>
                    <th>Quantity</th>
                    <th>Price</th>
                </tr>
                </thead>
                <tbody>
                {#key cartChangesCompleted}
                    {#each cartItems as item}
                        <tr>
                            <td><a href="/Products/{item.Product.ID}">{item.Product.Label}</a></td>
                            <td>
                                <div style="display: flex; flex-direction: row; align-items: center">
                                    <input class="qty" type="number" bind:value="{item.Quantity}" min=1 max={item.Product.Stock}>
                                    <div style="padding-left:10px"></div>
                                    <button on:click={() => del(`${item.Product.ID}`)} style="width: 25px; height:35px">X</button>
                                </div>
                            </td>
                            <td>{roundNumber(item.Product.Price * item.Quantity, 2)} €</td>
                        </tr>
                    {/each}
                {/key}
                </tbody>
            </table>
        </div>
    {:else}
        <h1>Your cart is empty!</h1>
        <div style="display: flex; flex-direction: row; align-items: center; justify-content: center">
            <h2>You may add items to your cart from the</h2>
            <div style="padding-left: 10px"></div>
            <h2><a href="/Products">product page</a></h2>
        </div>

        <div style="display: flex; flex-direction: row; align-items: center; justify-content: center">
            <h2>Or navigate all products by category on</h2>
            <div style="padding-left: 10px"></div>
            <h2><a href="/Categories">categories page</a></h2>
        </div>


    {/if}
{:else}
    <h1>You must be logged in to view your cart!</h1>
{/if}

<style>
    body {
        font-family: Arial, sans-serif;
        background-color: #f2f2f2;
    }

    h1 {
        text-align: center;
    }

    table {
        margin: 0 auto;
        border-collapse: collapse;
        width: auto;
        background-color: #fff;
        box-shadow: 0 0 20px rgba(0,0,0,0.15);
        table-layout: fixed;
        border: 1px solid #ddd;
        border-top: none;
        border-bottom: none;
    }

    tr {
        border-bottom: 1px solid #ddd;
    }

    th, td {
        padding: 12px 15px;
        text-align: left;
        border-bottom: 1px solid #ddd;
    }

    tr:hover {
        background-color: #f5f5f5;
    }

    th {
        background-color: #12171a;
        color: white;
        font-weight: bold;
    }

    .qty {
        width: 60px;
        height: 30px;
        padding: 0 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
        border-bottom: 3px solid #ddd;
        box-sizing: border-box;
        font-size: 16px;
    }

    .container {
        height: 100vh;
    }

    .container table {
        overflow-x: hidden;
    }

    td:last-child {
        border-bottom: none;
    }
</style>