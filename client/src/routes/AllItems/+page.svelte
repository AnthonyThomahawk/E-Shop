<script>
    import ItemFrame from './ItemFrame.svelte'
</script>
<h1>All items</h1>
<div style="display:flex;">
    <ItemFrame ItemName="First item" ItemPrice="999$" imagePath="./Images/svelte.svg"/>
    <div style="padding:10px"></div>
    <ItemFrame ItemName="Second item" ItemPrice="50$" imagePath="./Images/store-icon.svg"/>
    <div style="padding:10px"></div>
    <ItemFrame ItemName="Third item" ItemPrice="15$" imagePath="./Images/svelte.svg"/>
</div>
