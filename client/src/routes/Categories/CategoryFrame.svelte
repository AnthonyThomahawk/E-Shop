<script lang="ts">
    export let Name = "name";
    export let Description = "description";
    export let Color = "#000000";
    export let ID = 0;
</script>

<div style="border-color: {Color}">
    <h2 style="color:{Color}">{Name}</h2>
    <h3>{Description}</h3>
    <a style="position: absolute; bottom: 0; padding-bottom: 10px" href="/Categories/{ID}">Go to category</a>
</div>

<style>
    div {
        position: relative;
        flex-direction: column;
        display:flex;
        justify-content:center;
        align-items: center;
        border: 2px solid black;
        border-radius: 5%;
        padding: 2%;
        min-width: 15%;
        min-height: 250px;
        text-align: justify;
    }
</style>