<script lang="ts">
    import CategoryFrame from "./CategoryFrame.svelte";
    import {getProductCategories} from "../../lib/Products";
    import {onMount} from "svelte";

    interface ICategory {
        ID: number;
        Description: string;
        Label: string;
    }

    let categories : Array<ICategory> = [];

    async function getCategories() {
        const res = await getProductCategories();

        categories = res.map((item : any) => ({
            ID: item.ID,
            Description: item.Description,
            Label: item.Label
        }));
    }

    onMount(getCategories);
</script>

<h1>Product categories</h1>

<div style="display: flex; align-items: center;justify-content: center">
    {#each categories as category}
        <CategoryFrame Color="black" Name="{category.Label}" ID="{category.ID}" Description="{category.Description}"/>
        <div style="padding-right: 15px"></div>
    {/each}
</div>