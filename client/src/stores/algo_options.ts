import OptionsStore from './options';
import type { AlgoOptions as ModelAlgoOptions } from '../models/algo_options';
import { derived } from 'svelte/store';
import lastMoveStore from './last_move';
import { SHAPES } from '../models/algo_options';

export default derived([OptionsStore, lastMoveStore], ($Stores): ModelAlgoOptions => {
	const shape_neighboor = $Stores[0].proximity.shape == SHAPES.neighbour;

	return <ModelAlgoOptions>{
		time_out: $Stores[0].time_out,
		position: {
			x: $Stores[1].x,
			y: $Stores[1].y
		},
		depth: {
			max: $Stores[0].depth.max,
			min: $Stores[0].depth.min,
			pruning_percentage: $Stores[0].depth.pruning ? $Stores[0].depth.pruning_percentage : 0
		},
		width: {
			max: $Stores[0].width.max,
			pruning_percentage: $Stores[0].width.pruning ? $Stores[0].width.pruning_percentage : 0,
			multi_threading: $Stores[0].width.multi_threading
		},
		proximity: {
			radius: shape_neighboor ? 1 : $Stores[0].proximity.radius,
			threshold: shape_neighboor ? 1 : $Stores[0].proximity.threshold,
			shape: shape_neighboor ? SHAPES.square : $Stores[0].proximity.shape,
			evolution: $Stores[0].proximity.evolution
		},
		heuristics: {
			alignement: $Stores[0].heuristics.alignement,
			capture: $Stores[0].heuristics.capture,
			random: $Stores[0].heuristics.random
		},
		suspicion: {
			radius: $Stores[0].suspicion.active ? $Stores[0].suspicion.radius : 0
		}
	};
});
