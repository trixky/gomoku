import { writable } from 'svelte/store';
import type BoardModel from '../models/board';
import type CellModel from '../models/cell';

function generateBoard(): BoardModel {
	const cells: CellModel[][] = new Array(19);
	for (let i = 0; i < 19; i++) {
		cells[i] = [];
		for (let j = 0; j < 19; j++) {
			cells[i].push(<CellModel>{ player: 0 });
		}
	}

	return <BoardModel>{
		cells
	};
}

function createBoardStore() {
	const { subscribe, set, update } = writable(generateBoard());

	return {
		subscribe,
		reset: () => set(generateBoard()),
		addPiece: (player: 1 | 2, x: number, y: number) => {
			update((board) => {
				board.cells[y][x] = <CellModel>{
					player
				};

				return board;
			});
		}
	};
}

const boardStore = createBoardStore();

export default boardStore;
