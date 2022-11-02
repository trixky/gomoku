<!-- ========================= SCRIPT -->
<script lang="ts">
	import Goban from '../components/goban/goban.svelte';
	import TopGoban from '../components/goban/top.svelte';
	import BottomGoban from '../components/goban/bottom.svelte';

	let mode_ia = true;
	let turn: 1 | 2 = 1;
</script>

<!-- ========================= HTML -->

<div class="game">
	<TopGoban {mode_ia} />
	<Goban />
	<BottomGoban />
</div>
