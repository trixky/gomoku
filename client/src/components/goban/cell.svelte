<!-- ========================= SCRIPT -->
<script lang="ts">
	import type CellModel from '../../models/cell';
	import Piece from './piece.svelte';
	import OptionsStore from '../../stores/options';
	import HeuristicStatsStore from '../../stores/heuristic_stats';
	import { SHAPES } from '../../models/algo_options';

	export let x: number;
	export let y: number;
	export let cell: CellModel;
	export let handleLeftClick: (x: number, y: number) => void;
	export let proximity: number;
	export let heuristic: number;

	$: piece = cell.player != 0;

	$: pink =
		$OptionsStore.proximity.threshold === proximity ||
		$OptionsStore.proximity.shape === SHAPES.neighbour;

	$: heuristic_ratio =
		(heuristic - $HeuristicStatsStore.min) / ($HeuristicStatsStore.max - $HeuristicStatsStore.min);

	$: heuristic_opacity = heuristic_ratio * 0.3 + 0.1;
</script>

<!-- ========================= HTML -->
<div
	title="{cell.player > 0
		? `player: ${cell.player}`
		: ''}&#013;x: {x}&#013;y: {y}&#013;proximity: {proximity}&#013;heuristic: {heuristic}"
	class="piece-emplacement"
	style={$OptionsStore.proximity.show
		? `background-color: rgba(${pink ? 255 : 0}, 120, 255, ${Math.min(proximity * 0.09, 0.95)})`
		: $OptionsStore.heuristics.show
		? `background-color: rgba(${120 - heuristic_ratio * 120}, 255, 0, ${heuristic_opacity})`
		: ''}
	on:mousedown={() => handleLeftClick(x, y)}
>
	{#if piece}
		<Piece player={cell.player} />
	{/if}
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.piece-emplacement {
		@apply flex justify-center items-center w-[26px] h-[26px] max-w-[26px] max-h-[26px] transition-all duration-75;
	}

	.piece-emplacement:hover {
		border: 2px solid rgb(111, 169, 207);
		cursor: pointer;
	}
</style>
