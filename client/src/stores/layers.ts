import AnalyzerStore from './analyzer';
import type LayerModel from '../models/layer';
import { derived } from 'svelte/store';

export default derived(AnalyzerStore, ($AnalyzerStore): LayerModel[] => {
	const total = <LayerModel>{
		total: 0,
		cutted_by_max_width: 0,
		cutted_by_time_out: 0,
		pruned_in_depth: 0,
		pruned_in_width: 0,
		saved_by_min_depth: 0,
		selected: 0
	};

	$AnalyzerStore.layers.forEach((layer) => {
		total.total += layer.total;
		total.cutted_by_max_width += layer.cutted_by_max_width;
		total.cutted_by_time_out += layer.cutted_by_time_out;
		total.pruned_in_depth += layer.pruned_in_depth;
		total.pruned_in_width += layer.pruned_in_width;
		total.saved_by_min_depth += layer.saved_by_min_depth;
		total.selected += layer.selected;
	});

	const layers = [...$AnalyzerStore.layers, total];

	return layers;
});
