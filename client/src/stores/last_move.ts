import { writable } from 'svelte/store';
import type MoveModel from '../models/move';

function initMove(): MoveModel {
	return <MoveModel>{
		player: 2,
		x: 0,
		y: 0
	};
}

const g_history: MoveModel[] = [];

function createLastMoveStore() {
	const { subscribe, update, set } = writable(initMove());

	return {
		subscribe,
		push: (x: number, y: number) =>
			update((last_move) => {
				g_history.push(<MoveModel>{
					player: last_move.player,
					x: last_move.x,
					y: last_move.y
				});
				last_move.x = x;
				last_move.y = y;
				last_move.player = last_move.player == 1 ? 2 : 1;

				return last_move;
			}),
		reset: () => {
			g_history.splice(0, g_history.length);
			set(initMove());
		},
		undo: () => {
			const previous_last_move = g_history.pop();

			if (previous_last_move != undefined) {
				set(previous_last_move);
			}
		}
	};
}

const lastMoveStore = createLastMoveStore();

export default lastMoveStore;
