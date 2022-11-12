<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import GobanStore from '../../stores/goban';
	import StringGobanStore from '../../stores/string_goban';
	import AlgoOptionsStore from '../../stores/algo_options';
	import PlayersInfoStore from '../../stores/players_info';
	import ProximityGobanStore from '../../stores/proximity_goban';
	import Cell from './cell.svelte';
	import PostNext from '../../api/post.next';
	import PostCheck from '../../api/post.check';
	import LastMoveStore from '../../stores/last_move';
	import LoadingStore from '../../stores/loading';
	import TimeStore from '../../stores/time';
	import AnalyzerStore from '../../stores/analyzer';
	import PlayersInfosStore from '../../stores/players_info';
	import type NextResponseModel from '../../models/next_response';
	import type CheckResponseModel from '../../models/check_response';
	import RulesStore from '../../stores/rules';

	let rules = false;

	function checkWinByCaptures(): boolean {
		if ($PlayersInfoStore.player_2.captures >= 10) {
			alert('player 2 win by captures!');
			return true;
		} else if ($PlayersInfoStore.player_1.captures >= 10) {
			alert('player 1 win by captures!');
			return true;
		}

		return false;
	}

	interface checkWinByAlignementModel {
		1: {
			win: boolean;
			loophole: boolean;
		};
		2: {
			win: boolean;
			loophole: boolean;
		};
	}

	function checkWinByAlignement(): checkWinByAlignementModel {
		const response = <checkWinByAlignementModel>{
			1: {
				win: false,
				loophole: false
			},
			2: {
				win: false,
				loophole: false
			}
		};

		let current_player: 0 | 1 | 2 = 1;
		let opponent_player: 0 | 1 | 2 = 2;
		let horizontal_loophole = false,
			vertical_loophole = false,
			diagonal_loophole = false,
			reversed_diagonal_loophole = false;
		let horizontal_alignement = 0,
			vertical_alignement = 0,
			diagonal_alignement = 0,
			reversed_diagonal_alignement = 0;

		function reset() {
			horizontal_loophole = false;
			vertical_loophole = false;
			diagonal_loophole = false;
			horizontal_alignement = 1;
			vertical_alignement = 1;
			diagonal_alignement = 1;
			reversed_diagonal_alignement = 0;
		}

		for (let y = 0; y < 19; y++)
			for (let x = 0; x < 19; x++) {
				current_player = $GobanStore.cells[y][x].player;

				opponent_player = <0 | 1 | 2>((current_player % 2) + 1);
				reset();

				if (current_player != 0) {
					if (response[current_player].win && response[current_player].loophole) continue;
					for (let i = 1; i < 5; i++) {
						const xpi = x + i;
						const xmi = x - i;
						const ypi = y + i;

						if (x < 15 && $GobanStore.cells[y][xpi].player === current_player) {
							horizontal_alignement++;
						}
						if (y < 15 && $GobanStore.cells[ypi][x].player === current_player) {
							vertical_alignement++;
						}
						if (x < 15 && y < 15 && $GobanStore.cells[ypi][xpi].player === current_player) {
							diagonal_alignement++;
						}
						if (x > 3 && y < 15 && $GobanStore.cells[ypi][xmi].player === current_player) {
							reversed_diagonal_alignement++;
						}
					}

					if (horizontal_alignement === 5) {
						response[current_player].win = true;
						response[current_player].loophole =
							response[current_player].loophole || horizontal_loophole;
					} else if (vertical_alignement === 5) {
						response[current_player].win = true;
						response[current_player].loophole =
							response[current_player].loophole || vertical_loophole;
					} else if (diagonal_alignement === 5) {
						response[current_player].win = true;
						response[current_player].loophole =
							response[current_player].loophole || diagonal_loophole;
					} else if (reversed_diagonal_alignement === 5) {
						response[current_player].win = true;
						response[current_player].loophole =
							response[current_player].loophole || reversed_diagonal_loophole;
					}
				}
			}

		if (response[1].win || response[2].win) console.log(response);
		return response;
	}

	async function handleCellClick(x: number, y: number) {
		if (!$LoadingStore) {
			if ($GobanStore.cells[y][x].player === 0) {
				LoadingStore.switch(true);

				const current_last_move = $LastMoveStore;

				LastMoveStore.push(x, y);
				GobanStore.addPiece($LastMoveStore.player, x, y);

				PostCheck($StringGobanStore, $AlgoOptionsStore, $PlayersInfoStore)
					.then((response) => {
						const json_response: CheckResponseModel = JSON.parse(response);
						const concerned_player = !Boolean($LastMoveStore.player - 1);

						if (!json_response.DoubleThree) {
							PlayersInfoStore.addCaptures(concerned_player, json_response.NbCaptured);

							TimeStore.reset();
							GobanStore.playersFromString(json_response.goban);

							// checkWinByCaptures();
							// checkWinByAlignement();

							PostNext($StringGobanStore, $AlgoOptionsStore, $PlayersInfoStore)
								.then((response) => {
									const json_response: NextResponseModel = JSON.parse(response);

									GobanStore.playersFromString(json_response.goban);
									PlayersInfoStore.add(json_response.players_info);
									AnalyzerStore.set(json_response.analyzer);
									LastMoveStore.push(
										json_response.options.position.x,
										json_response.options.position.y
									);
									GobanStore.addPiece($LastMoveStore.player, $LastMoveStore.x, $LastMoveStore.y);
									GobanStore.heuristicFromString(json_response.heuristic_goban);
									TimeStore.set(json_response.options.time);

									// checkWinByCaptures();
									// checkWinByAlignement();

									LoadingStore.switch(false);
								})
								.catch(() => {
									alert('an error occured from api [next]');
									location.reload();
								});
						} else {
							alert("you can't play here");
							GobanStore.removePiece(x, y);
							LastMoveStore.push(current_last_move.x, current_last_move.y);
							LoadingStore.switch(false);
						}
					})
					.catch(() => {
						alert('an error occured from api [check]');
						location.reload();
					});
			}
		}
	}

	function handleKeyDown(e: any) {
		if (e.keyCode === 88) {
			// If 'x' is pressed
			PlayersInfosStore.reset();
			GobanStore.reset();
		}
	}

	function handleMouseEnterRules() {
		rules = true;
	}

	function handleMouseLeaveRules() {
		rules = false;
	}
</script>

<!-- ========================= HTML -->
<div class:rules class="goban-container">
	<div class:rules class="rules-container">
		<ul>
			{#each RulesStore as rule}
				<li>
					<p>{@html rule}</p>
				</li>
			{/each}
		</ul>
	</div>
	<div class:rules class="goban">
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
	<button class="rules" on:mouseenter={handleMouseEnterRules} on:mouseleave={handleMouseLeaveRules}
		>rules<img src="/info-icon.svg" alt="information icon" /></button
	>
</div>

<svelte:window on:keydown={handleKeyDown} />

<!-- ========================= CSS -->
<style lang="postcss">
	.goban-container {
		@apply relative h-fit w-fit m-auto mt-2 bg-white transition-all duration-300;
		border: 2px solid black;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}

	.goban-container.rules {
		border: 2px solid rgba(0, 0, 0, 0.03);
		box-shadow: rgba(0, 0, 0, 0.02) 0px 5px 15px;
	}

	.goban {
		@apply relative m-6 transition-all duration-300;
	}

	.goban.rules {
		@apply opacity-[3%];
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

	button.rules {
		@apply absolute inline -bottom-[39px] right-0 transition-all duration-300 opacity-60 hover:opacity-100;
	}

	button.rules > img {
		@apply inline w-[14px] h-[14px] ml-1;
	}

	.rules-container {
		@apply absolute top-0 left-0 p-8 opacity-0 transition-all duration-300;
	}

	.rules-container.rules {
		@apply opacity-100;
	}

	.rules-container > ul > li {
		@apply mb-3;
	}
</style>
