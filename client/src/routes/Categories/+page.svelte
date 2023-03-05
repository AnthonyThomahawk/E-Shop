<script lang="ts">
    import CategoryFrame from "./CategoryFrame.svelte";
</script>

<h1>Product categories</h1>
<div style="display: flex">
    <CategoryFrame Color="green" Name="Green" Description="Green teas are cool"/>
    <div style="padding-right: 10px">
    </div>
    <CategoryFrame Color="black" Name="Black" Description="Black teas are cool"/>
</div>