import { writable } from 'svelte/store';
import type { PlayersInfo as PlayersInfoModel } from '../models/players_info';

function initPlayersInfo(): PlayersInfoModel {
	return <PlayersInfoModel>{
		player_1: {
			pieces: 0,
			captures: 0
		},
		player_2: {
			pieces: 0,
			captures: 0
		}
	};
}

function createPlayersInfoStore() {
	const { subscribe, set, update } = writable(initPlayersInfo());

	return {
		subscribe,
		reset: () => {
			set(initPlayersInfo());
		},
		setCaptures: (player: boolean, nbr: number) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].captures = nbr;
				return players_info;
			}),
		setPieces: (player: boolean, nbr: number) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].pieces = nbr;
				return players_info;
			})
	};
}

const playersInfoStore = createPlayersInfoStore();

export default playersInfoStore;
