<!-- ========================= SCRIPT -->
<script lang="ts">
	import type CellModel from '../../models/cell';
	import Piece from './piece.svelte';

	export let x: number;
	export let y: number;
	export let cell: CellModel;
	export let debug_mode: boolean;
	export let handleLeftClick: (player: 1 | 2, x: number, y: number) => void;
</script>

<!-- ========================= HTML -->
<div
	class:debug={debug_mode}
	class="piece-emplacement"
	on:mousedown={() => handleLeftClick(cell.player === 0 || cell.player === 2 ? 1 : 2, x, y)}
>
	{#if cell.player != 0}
		<Piece player={cell.player} />
	{/if}
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.piece-emplacement {
		@apply flex justify-center items-center w-[26px] h-[26px] max-w-[26px] max-h-[26px];
	}

	.piece-emplacement.debug {
		border: 1px solid rgba(255, 0, 0, 0.151);
	}

	.piece-emplacement:hover.debug,
	.piece-emplacement:hover {
		border: 2px solid rgb(111, 169, 207);
		border-radius: 3px;
		cursor: pointer;
	}
</style>
