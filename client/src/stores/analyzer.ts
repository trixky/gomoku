import { writable } from 'svelte/store';
import type Analyzer from '../models/analyzer';

function createAnalyzerStore() {
	const { subscribe, set } = writable(<Analyzer>{
		layers: []
	});

	return {
		subscribe,
		reset: () =>
			set(<Analyzer>{
				layers: []
			}),
		set
	};
}

const analyzerStore = createAnalyzerStore();

export default analyzerStore;
