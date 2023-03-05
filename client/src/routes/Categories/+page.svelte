<script lang="ts">
    import CategoryFrame from "./CategoryFrame.svelte";
</script>

<h1>Product categories</h1>
<div style="display: flex; align-items: center;justify-content: center">
    <CategoryFrame Color="black" Name="Black" ID="1" Description="Discover unique black and black flavored tea varieties from Sri Lanka (Ceylon), China and India."/>
    <div style="padding-right: 15px"></div>
    <CategoryFrame Color="green" Name="Green" ID="2" Description="The best green and green flavored teas from Ceylon, China, Japan and Taiwan."/>
    <div style="padding-right: 15px"></div>
    <CategoryFrame Color="gray" Name="White" ID="3" Description="Light and full flavored, low in caffeine, white and white flavored tea is the best choice for anytime of the day."/>
    <div style="padding-right: 15px"></div>
</div>
<div style="padding-top: 15px"></div>
<div style="display: flex; align-items: center;justify-content: center">
    <CategoryFrame Color="brown" Name="Oolong" ID="4" Description="It balances between black and green tea. Either low or high fermented it is always a tasteful choice."/>
    <div style="padding-right: 15px"></div>
    <CategoryFrame Color="red" Name="Rooibos" ID="5" Description="The South African infusion. Full in Vitamin C, without any caffeine."/>
    <div style="padding-right: 15px"></div>
</div>