import { writable } from 'svelte/store';
import type PlayerInfoModel from '../models/player_info';

function initPlayerInfos(): PlayerInfoModel[] {
	return <PlayerInfoModel[]>[
		{
			pieces: 0,
			captures: 0
		},
		{
			pieces: 0,
			captures: 0
		}
	];
}

function createPlayerInfosStore() {
	const { subscribe, set, update } = writable(initPlayerInfos());

	return {
		subscribe,
		reset: () => {
			set(initPlayerInfos());
		},
		addCaptures: (player: boolean, nbr: number) =>
			update((player_infos) => {
				player_infos[Number(player)].captures += nbr;
				return player_infos;
			}),
		setPieces: (player: boolean, nbr: number) =>
			update((player_infos) => {
				player_infos[Number(player)].pieces = nbr;
				return player_infos;
			})
	};
}

const playerInfosStore = createPlayerInfosStore();

export default playerInfosStore;
