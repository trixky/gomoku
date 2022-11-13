import { writable } from 'svelte/store';
import type MoveModel from '../models/move';

function initMove(): MoveModel {
	return <MoveModel>{
		player: 2,
		x: 0,
		y: 0
	};
}

function createLastMoveStore() {
	const { subscribe, update, set } = writable(initMove());

	return {
		subscribe,
		push: (x: number, y: number) =>
			update((last_move) => {
				last_move.x = x;
				last_move.y = y;
				last_move.player = last_move.player == 1 ? 2 : 1;

				return last_move;
			}),
		reset: () => {
			set(initMove());
		}
	};
}

const lastMoveStore = createLastMoveStore();

export default lastMoveStore;
