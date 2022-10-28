import OptionsStore from './options';
import type { AlgoOptions as ModelAlgoOptions } from '../models/algo_options';
import { derived } from 'svelte/store';
import lastMoveStore from './last_move';
import Config from '../config';

export default derived([OptionsStore, lastMoveStore], ($Stores): ModelAlgoOptions => {
	return <ModelAlgoOptions>{
		timeout: 1000,
		position: {
			x: $Stores[1].x,
			y: $Stores[1].y
		},
		depth: {
			max: Config.options.depth.max.default,
			min: Config.options.depth.min.default,
			pruning: Config.options.depth.pruning.default
		},
		proximity: {
			radius: $Stores[0].proximity.radius,
			threshold: $Stores[0].proximity.threshold,
			shape: $Stores[0].proximity.shape,
			no_update: Config.options.proximity.no_update.default
		}
	};
});
