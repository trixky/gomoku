import { writable } from 'svelte/store';
import type MoveModel from '../models/move';

function createLastMoveStore() {
	const { subscribe, set } = writable(<MoveModel>{
		player: 1,
		x: 0,
		y: 0
	});

	return {
		subscribe,
		update: (player: 1 | 2, x: number, y: number) =>
			set(<MoveModel>{
				player,
				x,
				y
			})
	};
}

const lastMoveStore = createLastMoveStore();

export default lastMoveStore;
