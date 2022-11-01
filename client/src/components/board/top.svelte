<!-- ========================= SCRIPT -->
<script lang="ts">
	import { fade } from 'svelte/transition';
	import LoadingStore from '../../stores/loading';
	import TimeStore from '../../stores/time';

	const fade_parameters = { duration: 250 };

	export let mode_ia: boolean;
	export let turn: 1 | 2;

	let player_1 = {
		pieces_nbr: 0,
		captures: 0
	};

	let player_2 = {
		pieces_nbr: 0,
		captures: 0
	};
</script>

<!-- ========================= HTML -->
<div class="top-container">
	<div class="player-container left">
		<h2>
			<span> Player 1 </span>
			{#if mode_ia}
				<span class="opacity-20"> (You) </span>
			{/if}
		</h2>
		<div class="player-infos left">
			<p>pieces: {player_1.pieces_nbr}</p>
			<p>captures: {player_1.captures}</p>
		</div>
	</div>
	<div class="player-container right">
		<h2>
			<span> Player 2 </span>
			{#if mode_ia}
				<span class="opacity-20"> (AI) </span>
			{/if}
		</h2>
		<div class="player-infos right">
			<p>
				time:&nbsp;<span class="min-w-[74px] text-right"
					>{#if $LoadingStore}<span in:fade={fade_parameters}>loading</span>{:else}
						<span in:fade={fade_parameters}>{$TimeStore} ms</span>
					{/if}</span
				>
			</p>
			<p>pieces: {player_2.pieces_nbr}</p>
			<p>captures: {player_2.captures}</p>
		</div>
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.top-container {
		@apply flex justify-between mb-2;
	}

	.player-container {
		@apply p-1;
	}

	.player-infos {
		@apply flex pt-[9px];
	}

	.player-container.left > h2 {
		@apply text-left;
	}

	.player-container.right > h2 {
		@apply text-right;
	}

	.player-container.left > .player-infos > p {
		@apply mr-2;
	}

	.player-container.right > .player-infos > p {
		@apply ml-2;
	}

	.player-infos > p {
		@apply flex;
	}
</style>
