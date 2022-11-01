import { writable } from 'svelte/store';
import type CellModel from '../models/cell';
import type GobanModel from '../models/goban';

function generateBoard(): GobanModel {
	const cells: CellModel[][] = new Array(19);
	for (let i = 0; i < 19; i++) {
		cells[i] = [];
		for (let j = 0; j < 19; j++) {
			cells[i].push(<CellModel>{ player: 0, heuristic: NaN });
		}
	}

	return <GobanModel>{
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
		},
		heuristicFromString: (str: string) => {
			const cells = str.split(',').map((cell) => parseInt(cell));

			update((board) => {
				let i = 0;

				for (let y = 0; y < 19; y++)
					for (let x = 0; x < 19; x++) {
						board.cells[y][x].heuristic = cells[i];
						i++;
					}

				return board;
			});
		}
	};
}

const boardStore = createBoardStore();

export default boardStore;
