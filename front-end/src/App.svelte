<script>
    import Header from './lib/Header.svelte'
    import Footer from './lib/Footer.svelte'
</script>

<Header/>
<div class="centered">
    <h1>Main content</h1>
</div>
<Footer/>

<style>
    .centered {
        text-align: center;
        padding: 200px;
    }
</style>