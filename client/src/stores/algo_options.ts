import OptionsStore from './options';
import type { AlgoOptions as ModelAlgoOptions } from '../models/algo_options';
import { derived } from 'svelte/store';
import lastMoveStore from './last_move';

export default derived([OptionsStore, lastMoveStore], ($Stores): ModelAlgoOptions => {
	return <ModelAlgoOptions>{
		time_out: $Stores[0].time_out,
		position: {
			x: $Stores[1].x,
			y: $Stores[1].y
		},
		depth: {
			max: $Stores[0].depth.max,
			min: $Stores[0].depth.min,
			pruning: $Stores[0].depth.pruning,
			reduction: $Stores[0].depth.reduction
		},
		width: {
			max: $Stores[0].width.max,
			pruning: $Stores[0].width.pruning,
			multi_threading: $Stores[0].width.multi_threading
		},
		proximity: {
			radius: $Stores[0].proximity.radius,
			threshold: $Stores[0].proximity.threshold,
			shape: $Stores[0].proximity.shape,
			evolution: $Stores[0].proximity.evolution
		},
		heuristics: {
			potential_alignement: $Stores[0].heuristics.potential_alignement,
			alignement: $Stores[0].heuristics.alignement,
			potential_capture: $Stores[0].heuristics.potential_capture,
			capture: $Stores[0].heuristics.capture,
			random: $Stores[0].heuristics.random
		}
	};
});
