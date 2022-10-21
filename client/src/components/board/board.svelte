<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import { page } from '$app/stores';

	$: debug_mode = $page.url.hash === "#debug"
</script>

<!-- ========================= HTML -->
<div class="board-container">
	<div class="board">
		<div class:debug={debug_mode} class="piece-container">
			{#each new Array(Config.board.piece_number) as _}
			<div class:debug={debug_mode} class="piece-emplacement" />
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
		@apply h-fit;
		border: 4px solid black;
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
		@apply absolute -top-[13px] -left-[14px] grid w-fit;
		grid-template-columns: repeat(19, 1fr);
		grid-template-rows: repeat(19, 1fr);
	}

	.piece-container.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}

	.piece-emplacement {
		@apply w-[26px] h-[26px] max-w-[26px] max-h-[26px];
	}

	.piece-emplacement.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}

	.cell {
		@apply w-[26px] h-[26px] max-w-[26px] max-h-[26px];
		border: 1px solid black;
	}
</style>
