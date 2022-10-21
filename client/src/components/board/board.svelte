<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import { page } from '$app/stores';

	$: debug_mode = $page.url.hash === "#debug"
</script>

<!-- ========================= HTML -->
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

<!-- ========================= CSS -->
<style lang="postcss">
	.board {
		position: relative;
		margin: 30px;
	}

	.cell-container {
		top: 0;
		left: 0;
		display: grid;
		grid-template-columns: repeat(18, 1fr);
		grid-template-rows: repeat(18, 1fr);
		width: fit-content;
		border: 1px solid black;
	}

	.piece-container {
		position: absolute;
		top: -13px;
		left: -13px;
		display: grid;
		grid-template-columns: repeat(19, 1fr);
		grid-template-rows: repeat(19, 1fr);
		width: fit-content;
	}

	.piece-container.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}

	.piece-emplacement {
		width: 26px;
		height: 26px;
		max-width: 26px;
		max-height: 26px;
	}

	.piece-emplacement.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}


	.cell {
		width: 26px;
		height: 26px;
		max-width: 26px;
		max-height: 26px;
		border: 1px solid black;
	}
</style>
