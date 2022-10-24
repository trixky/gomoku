<!-- ========================= SCRIPT -->
<script lang="ts">
	export let player: 0 | 1 | 2;
</script>

<!-- ========================= HTML -->
<div class:white={player === 1} class:black={player === 2} class="piece" />

<!-- ========================= CSS -->
<style lang="postcss">
	.piece {
		@apply w-[22px] h-[22px] rounded-full;
		background-color: red;
		border: solid 2px blue;
	}

	.piece.white {
		background-color: #fff;
		border: solid 2px #555;
	}

	.piece.black {
		background-color: #555;
		border: solid 2px #111;
	}
</style>
