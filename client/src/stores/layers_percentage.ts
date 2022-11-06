import LayersStore from './layers';
import type LayerModel from '../models/layer';
import { derived } from 'svelte/store';

export default derived(LayersStore, ($LayersStore): LayerModel[] => {
	return $LayersStore.map((layer) => {
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
});
