<!-- ========================= SCRIPT -->
<script lang="ts">
	import { fade } from 'svelte/transition';
	import LastMoveStore from '../../stores/last_move';
	import LoadingStore from '../../stores/loading';
	import TimeStore from '../../stores/time';
	import GobanStore from '../../stores/goban';
	import PlayersInfosStore from '../../stores/players_info';
	import PieceNumberStore from '../../stores/piece_number';
	import { vsStore as VsStore, Modes as OpponentsModes } from '../../stores/vs';
	import WinStore from '../../stores/win';
	import Config from '../../config';

	function handleReset() {
		WinStore.reset();
		LastMoveStore.reset();
		PlayersInfosStore.reset();
		GobanStore.reset();
	}
</script>

<!-- ========================= HTML -->
<div class="top-container">
	<div class="player-container left">
		<h2>
			<span class:my-turn={$LastMoveStore.player === 2}> Player 1 </span>
			{#if $VsStore === OpponentsModes[0]}
				<span class="opacity-20" transition:fade={Config.animation.fade.slow}> (You) </span>
			{/if}
		</h2>
		<div class="player-infos left">
			<p class="pieces">pieces: {$PieceNumberStore.player_1}</p>
			<p class="captures">captures: {$PlayersInfosStore.player_1.captures}</p>
			<button class="reset" on:click={handleReset}
				>reset <img src="/reset-icon.svg" alt="reset icon" /></button
			>
		</div>
	</div>
	<div class="player-container right">
		<h2>
			{#if $VsStore === OpponentsModes[0]}
				<span class="opacity-20" transition:fade={Config.animation.fade.slow}> (AI) </span>
			{/if}
			<span class:my-turn={$LastMoveStore.player === 1}> Player 2 </span>
		</h2>
		<div class="player-infos right">
			{#if $VsStore === OpponentsModes[0]}
				<p transition:fade={Config.animation.fade.slow} class="time">
					time:&nbsp;<span class="min-w-[74px] text-right"
						>{#if $LoadingStore}<span in:fade={Config.animation.fade.slow} class="blink"
								>loading</span
							>{:else}
							<span in:fade={Config.animation.fade.slow}>{$TimeStore} ms</span>
						{/if}</span
					>
				</p>
			{/if}
			<p class="pieces">pieces: {$PieceNumberStore.player_2}</p>
			<p class="captures">captures: {$PlayersInfosStore.player_2.captures}</p>
		</div>
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.top-container {
		@apply relative flex justify-between mb-2;
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

	.player-container.right > .player-infos > p:not(.time) {
		@apply ml-2;
	}

	.player-infos > p {
		@apply flex;
	}

	.pieces {
		@apply w-[70px];
	}
	.captures {
		@apply w-[73px];
	}

	.reset {
		@apply ml-[10px] rounded-sm transition-all duration-300 opacity-60 hover:opacity-100;
	}

	.reset > img {
		@apply inline w-[14px] h-[14px] ml-1;
	}

	h2 > span {
		@apply transition-all duration-300 border-b-transparent border-b-[1px] pb-1;
	}

	h2 > .my-turn {
		@apply border-b-neutral-500;
	}

	button.new-game {
		@apply transition-all duration-300 opacity-60 hover:opacity-100;
	}
</style>
