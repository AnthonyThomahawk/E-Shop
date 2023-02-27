<script>
    import ItemFrame from './ItemFrame.svelte'
</script>
<h1>Products</h1>
<div style="display:flex;">
    <ItemFrame ItemName="First item" ItemPrice="999$" ItemID="1" imagePath="./Images/svelte.svg"/>
    <div style="padding:10px"></div>
    <ItemFrame ItemName="Second item" ItemPrice="50$" ItemID="2" imagePath="./Images/store-icon.svg"/>
    <div style="padding:10px"></div>
    <ItemFrame ItemName="Third item" ItemPrice="15$" ItemID="3" imagePath="./Images/svelte.svg"/>
</div>
