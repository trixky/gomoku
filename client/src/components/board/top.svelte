<!-- ========================= SCRIPT -->
<script lang="ts">

export let mode_ia: boolean;
export let turn: 1 | 2;

let player_1 = {
    pieces_nbr: 0,
    max: 0,
}

let player_2 = {
    pieces_nbr: 0,
    max: 0,
}

</script>

<!-- ========================= HTML -->
<div class="top-container">
    <div class="player-container left">
        <h2>
        <span class:turn={turn === 1}>
            Player 1
        </span>
            {#if mode_ia}
            <span class="opacity-20">
                (You)
            </span>
            {/if}
        </h2>
        <div class="player-infos left">
            <p>pieces: {player_1.pieces_nbr}</p>
            <p>max: {player_1.max}</p>
        </div>
    </div>
    <div class="player-container right">
        <h2>
            <span class:turn={turn === 2}>
                Player 2
            </span>
            {#if mode_ia}
            <span class="opacity-20">
                (AI)
            </span>
            {/if}
        </h2>
        <div class="player-infos right">
            <p>pieces: {player_2.pieces_nbr}</p>
            <p>max: {player_2.max}</p>
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
        @apply flex text-sm;
    }

    .turn {
        @apply underline underline-offset-2;
    }
</style>