import OptionsStore from './options';
import type { AlgoOptions as ModelAlgoOptions } from '../models/algo_options';
import { SHAPE_SQUARE } from '../models/algo_options';
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
			max: 2,
			min: 1,
			pruning: true
		},
		proximity: {
			radius: $Stores[0].proximity.radius,
			threshold: $Stores[0].proximity.threshold,
			shape: SHAPE_SQUARE,
			no_update: false
		}
	};
});
