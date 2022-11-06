import LayersStore from './layers';
import OptionsStore from './options';
import type LayerModel from '../models/layer';
import { derived } from 'svelte/store';

export default derived([LayersStore, OptionsStore], ($Store): LayerModel[] => {
	let layers_percentage = $Store[0].map((layer) => {
		const percentage_layer = <LayerModel>{
			cutted_by_max_width: layer.cutted_by_max_width,
			cutted_by_time_out: layer.cutted_by_time_out,
			pruned_in_depth: layer.pruned_in_depth,
			pruned_in_width: layer.pruned_in_width,
			saved_by_min_depth: layer.saved_by_min_depth,
			selected: layer.selected,
			total: layer.total
		};

		percentage_layer.cutted_by_max_width =
			percentage_layer.total != 0
				? parseFloat((percentage_layer.cutted_by_max_width / percentage_layer.total).toFixed(2))
				: percentage_layer.cutted_by_max_width == 0
				? 0
				: 100;
		percentage_layer.cutted_by_time_out =
			percentage_layer.total != 0
				? parseFloat(
						((percentage_layer.cutted_by_time_out * 100) / percentage_layer.total).toFixed(2)
				  )
				: percentage_layer.cutted_by_time_out == 0
				? 0
				: 100;
		percentage_layer.pruned_in_depth =
			percentage_layer.total != 0
				? parseFloat(((percentage_layer.pruned_in_depth * 100) / percentage_layer.total).toFixed(2))
				: percentage_layer.pruned_in_depth == 0
				? 0
				: 100;
		percentage_layer.pruned_in_width =
			percentage_layer.total != 0
				? parseFloat(((percentage_layer.pruned_in_width * 100) / percentage_layer.total).toFixed(2))
				: percentage_layer.pruned_in_width == 0
				? 0
				: 100;
		percentage_layer.saved_by_min_depth =
			percentage_layer.total != 0
				? parseFloat(
						((percentage_layer.saved_by_min_depth * 100) / percentage_layer.total).toFixed(2)
				  )
				: percentage_layer.saved_by_min_depth == 0
				? 0
				: 100;
		percentage_layer.selected =
			percentage_layer.total != 0
				? parseFloat(((percentage_layer.selected * 100) / percentage_layer.total).toFixed(2))
				: percentage_layer.selected == 0
				? 0
				: 100;

		return percentage_layer;
	});

	if ($Store[1].analyzer.rounded) {
		layers_percentage = layers_percentage.map((layer) => {
			layer.cutted_by_max_width = parseInt(layer.cutted_by_max_width.toFixed());
			layer.cutted_by_time_out = parseInt(layer.cutted_by_time_out.toFixed());
			layer.pruned_in_depth = parseInt(layer.pruned_in_depth.toFixed());
			layer.pruned_in_width = parseInt(layer.pruned_in_width.toFixed());
			layer.saved_by_min_depth = parseInt(layer.saved_by_min_depth.toFixed());
			layer.selected = parseInt(layer.selected.toFixed());

			return layer;
		});
	}

	return layers_percentage;
});
