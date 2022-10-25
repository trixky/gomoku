import OptionsStore from './options';
import type ModelAlgoOptions from '../models/algo_options';
import { derived } from 'svelte/store';

export default derived(OptionsStore, ($OptionsStore): ModelAlgoOptions => {
	return <ModelAlgoOptions>{
		proximity: {
			radius: $OptionsStore.proximity.radius,
			threshold: $OptionsStore.proximity.threshold
		}
	};
});
