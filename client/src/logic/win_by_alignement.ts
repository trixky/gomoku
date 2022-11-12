import type CellModel from '../models/cell';
import cellCanBeCaptured from './cell_can_be_captured';

interface checkWinByAlignementModel {
	1: {
		win: boolean;
		loophole: boolean;
	};
	2: {
		win: boolean;
		loophole: boolean;
	};
}

export default function winByAlignement(cells: CellModel[][]): checkWinByAlignementModel {
	const response = <checkWinByAlignementModel>{
		1: {
			win: false,
			loophole: false
		},
		2: {
			win: false,
			loophole: false
		}
	};

	let current_player: 0 | 1 | 2;
	let global_loophole = false,
		horizontal_loophole = false,
		vertical_loophole = false,
		diagonal_loophole = false,
		reversed_diagonal_loophole = false;
	let horizontal_alignement = 1,
		vertical_alignement = 1,
		diagonal_alignement = 1,
		reversed_diagonal_alignement = 1;

	function reset() {
		global_loophole = false;
		horizontal_loophole = false;
		vertical_loophole = false;
		diagonal_loophole = false;
		horizontal_alignement = 1;
		vertical_alignement = 1;
		diagonal_alignement = 1;
		reversed_diagonal_alignement = 1;
	}

	for (let y = 0; y < 19; y++)
		for (let x = 0; x < 19; x++) {
			current_player = cells[y][x].player;

			reset();

			if (current_player != 0) {
				if (response[current_player].win && response[current_player].loophole) continue;

				if (cellCanBeCaptured(x, y, cells)) {
					global_loophole = true;
				}

				for (let i = 1; i < 5; i++) {
					const xpi = x + i;
					const xmi = x - i;
					const ypi = y + i;

					if (x < 15 && cells[y][xpi].player === current_player) {
						// Horizontal
						//
						//    s . . . e
						//
						if (cellCanBeCaptured(xpi, y, cells)) {
							horizontal_loophole = true;
						}
						horizontal_alignement++;
					}
					if (y < 15 && cells[ypi][x].player === current_player) {
						// Vertical
						//
						//        s
						//        .
						//        .
						//        .
						//        e
						//
						if (cellCanBeCaptured(x, ypi, cells)) {
							vertical_loophole = true;
						}
						vertical_alignement++;
					}
					if (x < 15 && y < 15 && cells[ypi][xpi].player === current_player) {
						// Diagonal
						//
						//    s
						//      .
						//        .
						//          .
						//            e
						//
						if (cellCanBeCaptured(xpi, ypi, cells)) {
							diagonal_loophole = true;
						}
						diagonal_alignement++;
					}
					if (x > 3 && y < 15 && cells[ypi][xmi].player === current_player) {
						// Reversed Diagonal
						//
						//            s
						//          .
						//        .
						//      .
						//    e
						//
						if (cellCanBeCaptured(xmi, ypi, cells)) {
							reversed_diagonal_loophole = true;
						}
						reversed_diagonal_alignement++;
					}
				}

				if (horizontal_alignement === 5) {
					response[current_player].win = true;
					response[current_player].loophole = horizontal_loophole || global_loophole;
				} else if (vertical_alignement === 5) {
					response[current_player].win = true;
					response[current_player].loophole = vertical_loophole || global_loophole;
				} else if (diagonal_alignement === 5) {
					response[current_player].win = true;
					response[current_player].loophole = diagonal_loophole || global_loophole;
				} else if (reversed_diagonal_alignement === 5) {
					response[current_player].win = true;
					response[current_player].loophole = reversed_diagonal_loophole || global_loophole;
				}
			}
		}

	return response;
}
