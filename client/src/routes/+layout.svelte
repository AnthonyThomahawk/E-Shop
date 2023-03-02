<script lang="ts">
    import Header from '../lib/Header.svelte'
    import Footer from '../lib/Footer.svelte'
</script>

<Header/>
<div style="padding:100px; align-items: center; justify-content: center">
    <slot/>
</div>
<Footer/>