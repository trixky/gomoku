import BoardStore from './goban';
import OptionsStore from './options';
import { SHAPES } from '../models/algo_options';
import { derived } from 'svelte/store';

export default derived([OptionsStore, BoardStore], ($Stores): number[][] => {
	const proximity_cells = new Array(19).fill(undefined).map(() => new Array(19).fill(0));

	const shape_neighboor = $Stores[0].proximity.shape == SHAPES.neighbour;
	const radius = shape_neighboor ? 1 : $Stores[0].proximity.radius;
	const radius_plus = radius + 1;
	const selection_threshold = shape_neighboor ? 1 : $Stores[0].proximity.threshold;

	$Stores[1].cells.forEach((cells, y) => {
		cells.forEach((cell, x) => {
			if (cell.player > 0) {
				for (let local_y = -radius; local_y <= radius; local_y++) {
					const real_y = local_y + y;

					if (real_y >= 0 && real_y < 19)
						for (let local_x = -radius; local_x <= radius; local_x++) {
							const real_x = local_x + x;

							const diff = radius_plus - Math.max(Math.abs(local_x), Math.abs(local_y));

							if (real_x >= 0 && real_x < 19) {
								if (
									$Stores[0].proximity.shape === SHAPES.square ||
									Math.abs(local_x) === Math.abs(local_y) ||
									local_x === 0 ||
									local_y === 0
								) {
									proximity_cells[real_y][real_x] = Math.min(
										proximity_cells[real_y][real_x] + diff,
										selection_threshold
									);
								}
							}
						}
				}
			}
		});
	});

	return proximity_cells;
});
