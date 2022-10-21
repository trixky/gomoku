<!-- ========================= SCRIPT -->
<script lang="ts">
    import Header from '../components/layer/header.svelte'
    import Footer from '../components/layer/footer.svelte'
	import '../app.css';
</script>

<!-- ========================= HTML -->
<Header />
<main>
    <slot />
</main>
<Footer />

<!-- ========================= CSS -->
<style lang="postcss">
    main {
        display: flex;
        justify-content: center;
    }
</style>
