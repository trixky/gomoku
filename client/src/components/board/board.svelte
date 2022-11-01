<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import BoardStore from '../../stores/board';
	import StringBoardStore from '../../stores/string_board';
	import AlgoOptionsStore from '../../stores/algo_options';
	import ProximityBoardStore from '../../stores/proximity_board';
	import Cell from './cell.svelte';
	import PostNext from '../../api/post.next';
	import LastMoveStore from '../../stores/last_move';
	import LoadingStore from '../../stores/loading';
	import TimeStore from '../../stores/time';
	import type ResponseModel from '../../models/response';

	function handleCellClick(x: number, y: number) {
		if (!$LoadingStore) {
			if ($BoardStore.cells[y][x].player === 0) {
				LoadingStore.switch(true);
				TimeStore.reset();
				LastMoveStore.push(x, y);
				BoardStore.addPiece($LastMoveStore.player, x, y);
				PostNext($StringBoardStore, $AlgoOptionsStore)
					.then((response) => {
						const json_response: ResponseModel = JSON.parse(response);

						LastMoveStore.push(json_response.options.position.x, json_response.options.position.y);
						BoardStore.addPiece($LastMoveStore.player, $LastMoveStore.x, $LastMoveStore.y);
						BoardStore.heuristicFromString(json_response.heuristic_goban);
						TimeStore.set(json_response.options.time);
						LoadingStore.switch(false);
					})
					.catch(() => {
						alert('an error occured from api');
						LoadingStore.switch(false);
					});
			}
		}
	}
</script>

<!-- ========================= HTML -->
<div class="board-container">
	<div class="board">
		<div class="piece-container">
			{#each $BoardStore.cells as cells, y}
				{#each cells as cell, x}
					<Cell
						{x}
						{y}
						{cell}
						handleLeftClick={handleCellClick}
						proximity={$ProximityBoardStore[y][x]}
						heuristic={cell.heuristic}
					/>
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
		border: 2px solid black;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}

	.board {
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
