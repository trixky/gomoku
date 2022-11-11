import GobanStore from './goban';
import { derived } from 'svelte/store';
import type PieceNumberModel from 'src/models/piece_number';

export default derived(GobanStore, ($GobanStore): PieceNumberModel => {
	const piece_number = <PieceNumberModel>{
		player_1: 0,
		player_2: 0
	};

	$GobanStore.cells.forEach((line) =>
		line.forEach((cell) => {
			if (cell.player != 0) {
				if (cell.player === 1) piece_number.player_1++;
				else piece_number.player_2++;
			}
		})
	);

	return piece_number;
});
