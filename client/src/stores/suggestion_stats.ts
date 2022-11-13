import GobanStore from './goban';
import { derived } from 'svelte/store';
import type HeuristicStatsModel from 'src/models/heuristic_stats';

export default derived(GobanStore, ($GobanStore): HeuristicStatsModel => {
	const suggestion_stats = <HeuristicStatsModel>{
		max: -Infinity,
		min: Infinity,
		unique: true
	};

	$GobanStore.cells.forEach((line) =>
		line.forEach((cell) => {
			if (!isNaN(cell.suggestion)) {
				if (cell.suggestion > suggestion_stats.max) suggestion_stats.max = cell.suggestion;
				if (cell.suggestion < suggestion_stats.min) suggestion_stats.min = cell.suggestion;
			}
		})
	);

	suggestion_stats.unique = suggestion_stats.min === suggestion_stats.max;

	return suggestion_stats;
});
