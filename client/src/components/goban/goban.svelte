<!-- ========================= SCRIPT -->
<script lang="ts">
	import { fade } from 'svelte/transition';
	import Config from '../../config';
	import GobanStore from '../../stores/goban';
	import StringGobanStore from '../../stores/string_goban';
	import AlgoOptionsStore from '../../stores/algo_options';
	import { vsStore as VsStore, Modes } from '../../stores/vs';
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
	import winByAlignement from '../../logic/win_by_alignement';
	import winByCapture from '../../logic/win_by_capture';
	import RulesStore from '../../stores/rules';
	import WinStore from '../../stores/win';
	import playersInfoStore from '../../stores/players_info';

	let rules = false;
	let rubber = false;

	$: winner = $WinStore.player != 0 && !$WinStore.loophole;

	function check_win(): boolean {
		let capture_winner = winByCapture($PlayersInfoStore);
		if (capture_winner != 0) {
			WinStore.setCapture(capture_winner);
			return true;
		}

		let alignement_winner = winByAlignement($GobanStore.cells);
		if (
			(alignement_winner[1].win && $WinStore.player === 1) ||
			(alignement_winner[2].win && $WinStore.player === 2)
		) {
			WinStore.confirmAlignement();
			return true;
		} else if (alignement_winner[1].win) {
			WinStore.setAlignement(1, alignement_winner[1].loophole);
			if (!alignement_winner[1].loophole) {
				return true;
			}
		} else if (alignement_winner[2].win) {
			WinStore.setAlignement(2, alignement_winner[2].loophole);
			if (!alignement_winner[2].loophole) {
				return true;
			}
		}
		return false;
	}

	function askSuggestion() {
		PostNext($StringGobanStore, $AlgoOptionsStore, $PlayersInfoStore)
			.then((response) => {
				const json_response: NextResponseModel = JSON.parse(response);

				AnalyzerStore.set(json_response.analyzer);
				GobanStore.heuristicFromString(json_response.heuristic_goban, true);
				TimeStore.set(json_response.options.time);
			})
			.catch(() => {
				alert('an error occured from api [next]');
				location.reload();
			});
	}

	function undoAll() {
		playersInfoStore.undo();
		WinStore.undo();
		LastMoveStore.undo();
		GobanStore.undo();
	}

	async function handleCellClick(x: number, y: number) {
		if (!$LoadingStore && !winner) {
			if ($GobanStore.cells[y][x].player === 0) {
				LoadingStore.switch(true);

				const current_last_move = $LastMoveStore;

				LastMoveStore.push(x, y);
				GobanStore.addPiece($LastMoveStore.player, x, y, true);

				PostCheck($StringGobanStore, $AlgoOptionsStore, $PlayersInfoStore)
					.then((response) => {
						const json_response: CheckResponseModel = JSON.parse(response);
						const concerned_player = !Boolean($LastMoveStore.player - 1);

						if (!json_response.DoubleThree) {
							PlayersInfoStore.addCaptures(concerned_player, json_response.NbCaptured);

							TimeStore.reset();
							GobanStore.playersFromString(json_response.goban);

							if (check_win()) {
								LoadingStore.switch(false);
								return;
							}

							if ($VsStore === Modes[0])
								PostNext($StringGobanStore, $AlgoOptionsStore, $PlayersInfoStore)
									.then((response) => {
										const json_response: NextResponseModel = JSON.parse(response);

										GobanStore.playersFromString(json_response.goban, true);
										PlayersInfoStore.add(json_response.players_info);
										AnalyzerStore.set(json_response.analyzer);
										LastMoveStore.push(
											json_response.options.position.x,
											json_response.options.position.y
										);
										GobanStore.heuristicFromString(json_response.heuristic_goban, false);
										TimeStore.set(json_response.options.time);

										if (check_win()) {
											LoadingStore.switch(false);
											return;
										}

										askSuggestion();
										LoadingStore.switch(false);
									})
									.catch(() => {
										alert('an error occured from api [next]');
										location.reload();
									});
							else {
								askSuggestion();
								LoadingStore.switch(false);
							}
						} else {
							alert("you can't play here");
							undoAll();
							LoadingStore.switch(false);
						}
					})
					.catch(() => {
						alert('an error occured from api [check]');
						location.reload();
					});
			} else if (rubber) {
				GobanStore.removePiece(x, y);
			}
		}
	}

	function handleKeyDown(e: any) {
		if (e.keyCode === 85) {
			// If 'u' is pressed

			undoAll();
		}

		if (e.keyCode === 73) {
			// If 'i' is pressed

			rubber = !rubber;
		}

		if (e.keyCode === 88) {
			// If 'x' is pressed

			WinStore.reset();
			LastMoveStore.reset();
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

	function handleReset() {
		WinStore.reset();
		LastMoveStore.reset();
		PlayersInfosStore.reset();
		GobanStore.reset();
	}
</script>

<!-- ========================= HTML -->
{#if winner && !rules}
	<div transition:fade={Config.animation.fade.normal} class="win">
		<h2>Player {$WinStore.player.toString()} win by {$WinStore.methode} !</h2>
		<button class="new-game" on:click={handleReset}>new game</button>
	</div>
{/if}
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
	<div class:rules class:winner class="goban">
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
						suggestion={cell.suggestion}
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

	.goban.winner {
		@apply opacity-20;
	}

	.win {
		@apply absolute top-44 bg-white py-4 w-full text-center z-10;
		border: 2px solid black;
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
