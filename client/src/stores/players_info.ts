import { writable } from 'svelte/store';
import type { PlayersInfo as PlayersInfoModel } from '../models/players_info';

function initPlayersInfo(): PlayersInfoModel {
	return <PlayersInfoModel>{
		player_1: {
			pieces: 0,
			captures: 0,
			alignement: false,
			win: false
		},
		player_2: {
			pieces: 0,
			captures: 0,
			alignement: false,
			win: false
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
		add: (players_info: PlayersInfoModel) => {
			update((_players_info) => {
				_players_info.player_1.captures += players_info.player_1.captures;
				_players_info.player_2.captures += players_info.player_2.captures;

				return players_info;
			});
		},
		setCaptures: (player: boolean, nbr: number) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].captures = nbr;
				return players_info;
			}),
		addCaptures: (player: boolean, nbr: number) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].captures += nbr;
				return players_info;
			}),
		setAlignement: (player: boolean, alignement: boolean) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].alignement = alignement;
				return players_info;
			}),
		setPieces: (player: boolean, nbr: number) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].pieces = nbr;
				return players_info;
			}),
		setWin: (player: boolean, win: boolean) =>
			update((players_info) => {
				players_info[player ? 'player_1' : 'player_2'].win = win;
				return players_info;
			})
	};
}

const playersInfoStore = createPlayersInfoStore();

export default playersInfoStore;
