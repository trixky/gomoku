<!-- ========================= SCRIPT -->
<script lang="ts">
	import Board from '../components/board/board.svelte';
	import TopBoard from '../components/board/top.svelte';
	import BottomBoard from '../components/board/bottom.svelte';

	let mode_ia = true;
	let turn: 1 | 2 = 1;
</script>

<!-- ========================= HTML -->

<div class="game">
	<TopBoard {mode_ia} {turn} />
	<Board />
	<BottomBoard />
</div>
