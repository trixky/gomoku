<!-- ========================= SCRIPT -->
<script lang="ts">
	import Goban from '../components/goban/goban.svelte';
	import TopGoban from '../components/goban/top.svelte';
	import BottomGoban from '../components/goban/bottom.svelte';
</script>

<!-- ========================= HTML -->

<div class="game">
	<TopGoban />
	<Goban />
	<BottomGoban />
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.game {
		@apply relative;
	}
</style>
