import type OptionsModel from '../models/options';
import Config from '../config';

export default function generateOptions(): OptionsModel {
	return <OptionsModel>{
		time_out: Config.options.time_out.default,
		depth: {
			max: Config.options.depth.max.default,
			min: Config.options.depth.min.default,
			pruning: Config.options.depth.pruning.default,
			pruning_percentage: Config.options.depth.pruning_percentage.default
		},
		width: {
			max: Config.options.width.max.default,
			pruning: Config.options.width.pruning.default,
			pruning_percentage: Config.options.depth.pruning_percentage.default,
			multi_threading: Config.options.width.multi_threading.default
		},
		proximity: {
			radius: Config.options.proximity.radius.default,
			threshold: Config.options.proximity.threshold.default,
			show: Config.options.proximity.show.default,
			evolution: Config.options.proximity.evolution.default,
			shape: Config.options.proximity.shape.default
		},
		heuristics: {
			alignement: Config.options.heuristics.values.default,
			capture: Config.options.heuristics.values.default,
			random: Config.options.heuristics.values.default,
			show: Config.options.heuristics.show.default
		},
		suggestion: {
			show: Config.options.suggestion.show.default
		},
		suspicion: {
			active: Config.options.suspicion.active.default,
			radius: Config.options.suspicion.radius.default
		},
		analyzer: {
			layered: Config.options.analyzer.layered.default,
			percentage: Config.options.analyzer.percentage.default,
			rounded: Config.options.analyzer.rounded.default
		}
	};
}
