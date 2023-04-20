<script lang="ts">
    export let Name = "name";
    export let Description = "description";
    export let Color = "#000000";
    export let ID = 0;
</script>

<div style="border-color: {Color}">
    <h2 style="color:{Color}">{Name}</h2>
    <h3>{Description}</h3>
    <a href="/Categories/{ID}">Go to category</a>
</div>

<style>
    div {
        flex-direction: column;
        display:flex;
        justify-content: initial;
        align-items: center;
        border: 2px solid black;
        border-radius:2%;
        width:250px;
        height:250px;
        padding:10px;
        text-align: justify;
    }
</style>