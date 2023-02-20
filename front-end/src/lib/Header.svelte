<script>
    export let title = "E-Shop";
    import Icon from '../assets/store-icon.svg'
</script>

<header>
    <img src={Icon} alt="Main icon">
    <h1>{title}</h1>
</header>

<style>
    img {
        height:50px;
        width:50px;
        padding-right:10px;
    }
    header {
        font-size: 75%;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        background-color: #333;
        color: #fff;
        height: 64px;
        padding: 0 16px;
    }

    h1 {
        margin: 0;
    }
</style>
