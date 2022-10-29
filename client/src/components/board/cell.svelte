<!-- ========================= SCRIPT -->
<script lang="ts">
	import type CellModel from '../../models/cell';
	import Piece from './piece.svelte';
	import OptionsStore from '../../stores/options';

	export let x: number;
	export let y: number;
	export let cell: CellModel;
	export let handleLeftClick: (x: number, y: number) => void;
	export let proximity: number;
</script>

<!-- ========================= HTML -->
<div
	title="{cell.player > 0
		? `player: ${cell.player}`
		: ''}&#013;x: {x}&#013;y: {y}&#013;proximity: {proximity}"
	class="piece-emplacement"
	style={$OptionsStore.proximity.show
		? `background-color: rgba(${
				$OptionsStore.proximity.threshold === proximity ? 255 : 0
		  }, 120, 255, ${Math.min(proximity * 0.09, 0.95)})`
		: ''}
	on:mousedown={() => handleLeftClick(x, y)}
>
	{#if cell.player != 0}
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
