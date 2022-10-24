<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import { page } from '$app/stores';
	import BoardStore from '../../stores/board';
	import Cell from './cell.svelte';

	$: debug_mode = $page.url.hash === '#debug';

	function handleCellClick(player: 1 | 2, x: number, y: number) {
		BoardStore.refreshPiece(player, x, y);
	}
</script>

<!-- ========================= HTML -->
<div class="board-container">
	<div class="board">
		<div class:debug={debug_mode} class="piece-container">
			{#each $BoardStore.cells as cells, y}
				{#each cells as cell, x}
					<Cell {x} {y} {cell} {debug_mode} handleLeftClick={handleCellClick} />
				{/each}
			{/each}
		</div>
		<div class="cell-container">
			{#each new Array(Config.board.cell_number) as _}
				<div class="cell" />
			{/each}
		</div>
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.board-container {
		@apply h-fit bg-white;
		border: 3px solid black;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}

	.board {
		@apply relative m-6;
	}

	.cell-container {
		@apply grid w-fit;
		grid-template-columns: repeat(18, 1fr);
		grid-template-rows: repeat(18, 1fr);
		border: 1px solid black;
	}

	.piece-container {
		@apply absolute -top-[13px] -left-[13px] grid w-fit;
		grid-template-columns: repeat(19, 1fr);
		grid-template-rows: repeat(19, 1fr);
	}

	.piece-container.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}

	.cell {
		@apply w-[26px] h-[26px] max-w-[26px] max-h-[26px];
		border: 1px solid black;
	}
</style>
