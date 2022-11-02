import type OptionsModel from '../models/options';
import Config from '../config';

export default function generateOptions(): OptionsModel {
	return <OptionsModel>{
		time_out: Config.options.time_out.default,
		depth: {
			max: Config.options.depth.max.default,
			min: Config.options.depth.min.default,
			pruning: Config.options.depth.pruning.default
		},
		width: {
			max: Config.options.width.max.default,
			pruning: Config.options.width.pruning.default,
			multi_threading: Config.options.width.multi_threading.default
		},
		proximity: {
			radius: Config.options.proximity.radius.default,
			threshold: Config.options.proximity.threshold.default,
			show: Config.options.proximity.show.default,
			evolution: Config.options.proximity.evolution.default,
			shape: Config.options.proximity.shape.default,
			reduction: Config.options.proximity.reduction.default
		},
		heuristics: {
			potential_alignement: Config.options.heuristics.values.default,
			alignement: Config.options.heuristics.values.default,
			potential_capture: Config.options.heuristics.values.default,
			capture: Config.options.heuristics.values.default,
			random: Config.options.heuristics.values.default,
			show: Config.options.heuristics.show.default
		}
	};
}