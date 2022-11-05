import AnalyzerStore from './analyzer';
import OptionsStore from './options';
import type LayerModel from '../models/layer';
import { derived } from 'svelte/store';

export default derived([AnalyzerStore, OptionsStore], ($Stores): LayerModel[] => {
	const last_layer_index = $Stores[0].layers.length - 1;

	const total = <LayerModel>{
		cutted_by_max_width: 0,
		cutted_by_time_out: 0,
		examined: 0,
		pruned_in_depth: 0,
		pruned_in_width: 0,
		saved_by_min_depth: 0,
		selected: 0
	};

	let layers = $Stores[0].layers.map((layer, index) => {
		const child_layer = index === last_layer_index ? layer : $Stores[0].layers[index + 1];

		const examined =
			layer.cutted_by_max_width +
			layer.cutted_by_time_out +
			layer.pruned_in_depth +
			layer.pruned_in_width +
			child_layer.selected;

		total.cutted_by_max_width += layer.cutted_by_max_width;
		total.cutted_by_time_out += layer.cutted_by_time_out;
		total.pruned_in_depth += layer.pruned_in_depth;
		total.pruned_in_width += layer.pruned_in_width;
		total.saved_by_min_depth += layer.saved_by_min_depth;
		total.selected += layer.selected;
		total.examined += examined;

		return <LayerModel>{
			...layer,
			examined
		};
	});

	layers.push(total);

	if ($Stores[1].analyzer.percentage) {
		layers = layers.map((layer) => {
			layer.cutted_by_max_width = parseFloat(
				(layer.cutted_by_max_width / layer.examined).toFixed(2)
			);
			layer.cutted_by_time_out = parseFloat(
				((layer.cutted_by_time_out * 100) / layer.examined).toFixed(2)
			);
			layer.pruned_in_depth = parseFloat(
				((layer.pruned_in_depth * 100) / layer.examined).toFixed(2)
			);
			layer.pruned_in_width = parseFloat(
				((layer.pruned_in_width * 100) / layer.examined).toFixed(2)
			);
			layer.saved_by_min_depth = parseFloat(
				((layer.saved_by_min_depth * 100) / layer.examined).toFixed(2)
			);
			layer.selected = parseFloat(((layer.selected * 100) / layer.examined).toFixed(2));

			return layer;
		});
	}

	return layers;
});
