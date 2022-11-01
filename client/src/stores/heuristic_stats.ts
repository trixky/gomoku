import GobanStore from './goban';
import { derived } from 'svelte/store';
import type HeuristicStatsModel from 'src/models/heuristic_stats';

export default derived(GobanStore, ($GobanStore): HeuristicStatsModel => {
	const heuristic_stats = <HeuristicStatsModel>{
		max: -Infinity,
		min: Infinity
	};

	$GobanStore.cells.forEach((line) =>
		line.forEach((cell) => {
			if (!isNaN(cell.heuristic)) {
				if (cell.heuristic > heuristic_stats.max) heuristic_stats.max = cell.heuristic;
				if (cell.heuristic < heuristic_stats.min) heuristic_stats.min = cell.heuristic;
			}
		})
	);

	return heuristic_stats;
});
