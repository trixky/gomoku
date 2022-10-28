import OptionsStore from './options';
import type { AlgoOptions as ModelAlgoOptions } from '../models/algo_options';
import { derived } from 'svelte/store';
import lastMoveStore from './last_move';

export default derived([OptionsStore, lastMoveStore], ($Stores): ModelAlgoOptions => {
	return <ModelAlgoOptions>{
		timeout: 1000,
		position: {
			x: $Stores[1].x,
			y: $Stores[1].y
		},
		depth: {
			max: $Stores[0].depth.max,
			min: $Stores[0].depth.min,
			pruning: $Stores[0].depth.pruning
		},
		proximity: {
			radius: $Stores[0].proximity.radius,
			threshold: $Stores[0].proximity.threshold,
			shape: $Stores[0].proximity.shape,
			evolution: $Stores[0].proximity.evolution
		}
	};
});
