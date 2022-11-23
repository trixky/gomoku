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

const g_history: GobanModel[] = [];

function deepBoardCopy(board: GobanModel) {
	return <GobanModel>{
		cells: board.cells.map((line) =>
			line.map(
				(cell) =>
					<CellModel>{ heuristic: cell.heuristic, player: cell.player, suggestion: cell.suggestion }
			)
		)
	};
}

function createBoardStore() {
	const { subscribe, set, update } = writable(generateBoard());

	return {
		subscribe,
		reset: () => {
			g_history.splice(0, g_history.length);
			set(generateBoard());
		},
		addPiece: (player: 1 | 2, x: number, y: number, history = false) => {
			update((board) => {
				if (history) g_history.push(deepBoardCopy(board));
				board.cells[y][x] = <CellModel>{
					player
				};

				return board;
			});
		},
		removePiece: (x: number, y: number) => {
			update((board) => {
				board.cells[y][x] = <CellModel>{
					player: 0
				};

				return board;
			});
		},
		playersFromString: (str: string, history = false) => {
			const cells = str.split('').map((cell) => <0 | 1 | 2>parseInt(cell));

			update((board) => {
				if (history) g_history.push(deepBoardCopy(board));

				let i = 0;

				for (let y = 0; y < 19; y++)
					for (let x = 0; x < 19; x++) {
						board.cells[y][x].player = cells[i];
						i++;
					}

				return board;
			});
		},
		heuristicFromString: (str: string, suggestion = false) => {
			const cells = str.split(',').map((cell) => parseInt(cell));

			update((board) => {
				let i = 0;

				for (let y = 0; y < 19; y++)
					for (let x = 0; x < 19; x++) {
						if (suggestion) board.cells[y][x].suggestion = cells[i];
						else board.cells[y][x].heuristic = cells[i];
						i++;
					}

				return board;
			});
		},
		undo: () => {
			const previous_board = g_history.pop();

			if (previous_board != undefined) {
				set(previous_board);
			}
		}
	};
}

const boardStore = createBoardStore();

export default boardStore;
