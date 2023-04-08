<script lang="ts">
    import CategoryFrame from "./CategoryFrame.svelte";
    import {getProductCategories} from "../../lib/Products";
    import {onMount} from "svelte";
    import { page } from '$app/stores';
    import {getLocalStorage, setLocalStorage} from "../../lib/LocalStorage";

    interface ICategory {
        ID: number;
        Description: string;
        Label: string;
        ColorHex: string;
    }

    let categories : Array<ICategory> = [];
    let currentPath = "";
    let previousPath;

    async function getCategories() {
        currentPath = $page.url.pathname;
        setLocalStorage('previousPath', currentPath);

        const [res, hd] = await getProductCategories();

        categories = res.map((item : any) => ({
            ID: item.ID,
            Description: item.Description,
            Label: item.Label,
            ColorHex: item.Color
        }));
    }

    onMount(getCategories);
</script>

<h1>Product categories</h1>

<div style="display: flex; align-items: center;justify-content: center">
    {#each categories as category}
        <CategoryFrame Color="{category.ColorHex}" Name="{category.Label}" ID="{category.ID}" Description="{category.Description}"/>
        <div style="padding-right: 15px"></div>
    {/each}
</div>