<script>
    export let title = "E-Shop";
    import Icon from '../assets/store-icon.svg'
</script>

<header>
    <img src={Icon} alt="Main icon">
    <h1>{title}</h1>
    <nav style="padding-left:20px;">
        <a href="/" style="color:#fff">home</a>
        <a href="/about" style="color:#fff">about</a>
    </nav>
</header>

<style>
    img {
        height:50px;
        width:50px;
        padding-right:10px;
    }
    header {
        font-size: 165%;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        background-color: #333;
        color: #fff;
        height: 100px;
        padding: 0 16px;
    }

    h1 {
        margin: 0;
    }
</style>
