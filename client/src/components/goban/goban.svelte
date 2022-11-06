<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import GobanStore from '../../stores/goban';
	import StringGobanStore from '../../stores/string_goban';
	import AlgoOptionsStore from '../../stores/algo_options';
	import ProximityGobanStore from '../../stores/proximity_goban';
	import Cell from './cell.svelte';
	import PostNext from '../../api/post.next';
	import LastMoveStore from '../../stores/last_move';
	import LoadingStore from '../../stores/loading';
	import TimeStore from '../../stores/time';
	import AnalyzerStore from '../../stores/analyzer';
	import type ResponseModel from '../../models/response';

	function handleCellClick(x: number, y: number) {
		if (!$LoadingStore) {
			if ($GobanStore.cells[y][x].player === 0) {
				LoadingStore.switch(true);
				TimeStore.reset();
				LastMoveStore.push(x, y);
				GobanStore.addPiece($LastMoveStore.player, x, y);
				PostNext($StringGobanStore, $AlgoOptionsStore)
					.then((response) => {
						const json_response: ResponseModel = JSON.parse(response);

						AnalyzerStore.set(json_response.analyzer);
						LastMoveStore.push(json_response.options.position.x, json_response.options.position.y);
						GobanStore.addPiece($LastMoveStore.player, $LastMoveStore.x, $LastMoveStore.y);
						GobanStore.heuristicFromString(json_response.heuristic_goban);
						TimeStore.set(json_response.options.time);
						LoadingStore.switch(false);
					})
					.catch(() => {
						alert('an error occured from api');
						LoadingStore.switch(false);
						location.reload();
					});
			}
		}
	}

	function handleKeyDown(e: any) {
		if (e.keyCode === 88) {
			// If 'x' is pressed

			// Reset the goban
			GobanStore.reset();
		}
	}
</script>

<!-- ========================= HTML -->
<div class="goban-container">
	<div class="goban">
		<div class="piece-container">
			{#each $GobanStore.cells as cells, y}
				{#each cells as cell, x}
					<Cell
						{x}
						{y}
						{cell}
						handleLeftClick={handleCellClick}
						proximity={$ProximityGobanStore[y][x]}
						heuristic={cell.heuristic}
					/>
				{/each}
			{/each}
		</div>
		<div class="cell-container">
			{#each new Array(Config.goban.cell_number) as _}
				<div class="cell" />
			{/each}
		</div>
	</div>
</div>

<svelte:window on:keydown={handleKeyDown} />

<!-- ========================= CSS -->
<style lang="postcss">
	.goban-container {
		@apply h-fit w-fit m-auto bg-white;
		border: 2px solid black;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}

	.goban {
		@apply relative m-6;
	}

	.cell-container {
		@apply grid w-fit;
		grid-template-columns: repeat(18, 1fr);
		grid-template-rows: repeat(18, 1fr);
		border: 0.5px solid black;
	}

	.piece-container {
		@apply absolute -top-[13px] -left-[13px] grid w-fit;
		grid-template-columns: repeat(19, 1fr);
		grid-template-rows: repeat(19, 1fr);
	}

	.cell {
		@apply w-[26px] h-[26px] max-w-[26px] max-h-[26px];
		border: 0.5px solid black;
	}
</style>
