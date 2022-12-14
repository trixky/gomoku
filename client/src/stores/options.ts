import Config from '../config';
import { writable } from 'svelte/store';
import { SHAPES } from '../models/algo_options';
import type OptionsModel from '..//models/options';
import generateOptions from '../../src/utils/generate_options';

function createOptionsStore() {
	const { subscribe, set, update } = writable(generateOptions());

	return {
		subscribe,
		reset: () => set(generateOptions()),
		// ----------------------- set ai
		setAi: (ai: OptionsModel) => {
			update((_ai) => {
				ai.proximity.show = _ai.proximity.show;
				ai.heuristics.show = _ai.heuristics.show;
				ai.suggestion.show = _ai.suggestion.show;
				ai.analyzer.layered = _ai.analyzer.layered;
				ai.analyzer.percentage = _ai.analyzer.percentage;
				ai.analyzer.rounded = _ai.analyzer.rounded;

				return ai;
			});
		},
		// ----------------------- set time out
		setTimeOut: (time_out: string) =>
			update((options) => {
				const _time_out = parseInt(time_out);

				if (
					!isNaN(_time_out) &&
					_time_out >= Config.options.time_out.min &&
					_time_out <= Config.options.time_out.max
				)
					options.time_out = _time_out;
				return options;
			}),
		// ----------------------- set proximity
		showProximity: () =>
			update((options) => {
				options.proximity.show = true;
				return options;
			}),
		hideProximity: () =>
			update((options) => {
				options.proximity.show = false;
				return options;
			}),
		setProximityRadius: (radius: string) =>
			update((options) => {
				const _radius = parseInt(radius);

				if (
					!isNaN(_radius) &&
					_radius >= Config.options.proximity.radius.min &&
					_radius <= Config.options.proximity.radius.max
				)
					options.proximity.radius = _radius;
				return options;
			}),
		setProximityThreshold: (threshold: string) =>
			update((options) => {
				const _threshold = parseInt(threshold);

				if (
					!isNaN(_threshold) &&
					_threshold >= Config.options.proximity.threshold.min &&
					_threshold <= Config.options.proximity.threshold.max
				)
					options.proximity.threshold = _threshold;

				return options;
			}),
		setProximityShape: (shape: string) =>
			update((options) => {
				const _shape = parseInt(shape);

				if (!isNaN(_shape) && Object.values(SHAPES).includes(_shape))
					options.proximity.shape = _shape;

				return options;
			}),
		setProximityEvolution: (evolution: boolean) =>
			update((options) => {
				options.proximity.evolution = evolution;
				return options;
			}),
		// ----------------------- set depth
		setDepthMax: (max: string) =>
			update((options) => {
				const _max = parseInt(max);

				if (
					!isNaN(_max) &&
					_max >= options.depth.min &&
					_max >= Config.options.depth.max.min &&
					_max <= Config.options.depth.max.max
				)
					options.depth.max = _max;

				return options;
			}),
		setDepthMin: (min: string) =>
			update((options) => {
				const _min = parseInt(min);

				if (
					!isNaN(_min) &&
					_min <= options.depth.max &&
					_min >= Config.options.depth.min.min &&
					_min <= Config.options.depth.min.max
				)
					options.depth.min = _min;

				return options;
			}),
		setDepthPruning: (pruning: boolean) =>
			update((options) => {
				options.depth.pruning = pruning;
				return options;
			}),
		setDepthPruningPercentage: (pruning_percentage: string) =>
			update((options) => {
				const _pruning_percentage = parseInt(pruning_percentage);

				if (
					!isNaN(_pruning_percentage) &&
					_pruning_percentage >= Config.options.depth.pruning_percentage.min &&
					_pruning_percentage <= Config.options.depth.pruning_percentage.max
				)
					options.depth.pruning_percentage = _pruning_percentage;

				return options;
			}),
		// ----------------------- set width
		setWidthPruning: (pruning: boolean) =>
			update((options) => {
				options.width.pruning = pruning;
				return options;
			}),
		setWidthPruningPercentage: (pruning_percentage: string) =>
			update((options) => {
				const _pruning_percentage = parseInt(pruning_percentage);

				if (
					!isNaN(_pruning_percentage) &&
					_pruning_percentage >= Config.options.width.pruning_percentage.min &&
					_pruning_percentage <= Config.options.width.pruning_percentage.max
				)
					options.width.pruning_percentage = _pruning_percentage;

				return options;
			}),
		setWidthMultiThreading: (multi_threading: boolean) =>
			update((options) => {
				options.width.multi_threading = multi_threading;
				return options;
			}),
		setWidthMax: (max: string) =>
			update((options) => {
				const _max = parseInt(max);

				if (
					!isNaN(_max) &&
					_max >= Config.options.width.max.min &&
					_max <= Config.options.width.max.max
				)
					options.width.max = _max;

				return options;
			}),
		// ----------------------- set heuristics
		showHeuristics: () =>
			update((options) => {
				options.heuristics.show = true;
				return options;
			}),
		hideHeuristics: () =>
			update((options) => {
				options.heuristics.show = false;
				return options;
			}),
		setHeuristicsAggro: (aggro: string) =>
			update((options) => {
				const _aggro = parseInt(aggro);

				if (!isNaN(_aggro)) options.heuristics.aggro = _aggro;
				return options;
			}),
		setHeuristicsDepthDivisor: (depth_divisor: string) =>
			update((options) => {
				const _depth_divisor = parseInt(depth_divisor);

				if (!isNaN(_depth_divisor)) options.heuristics.depth_divisor = _depth_divisor;
				return options;
			}),
		setHeuristicsAlignement: (alignement: string) =>
			update((options) => {
				const _alignement = parseInt(alignement);

				if (!isNaN(_alignement)) options.heuristics.alignement = _alignement;
				return options;
			}),
		setHeuristicsCapture: (capture: string) =>
			update((options) => {
				const _capture = parseInt(capture);

				if (!isNaN(_capture)) options.heuristics.capture = _capture;
				return options;
			}),
		setHeuristicsRandom: (random: string) =>
			update((options) => {
				const _random = parseInt(random);

				if (!isNaN(_random)) options.heuristics.random = _random;
				return options;
			}),
		// ----------------------- set suggestion
		showSuggestion: () =>
			update((options) => {
				options.suggestion.show = true;
				return options;
			}),
		hideSuggestion: () =>
			update((options) => {
				options.suggestion.show = false;
				return options;
			}),
		// ----------------------- set suspicion
		activeSuspicion: () =>
			update((options) => {
				options.suspicion.active = true;
				return options;
			}),
		disableSuspicion: () =>
			update((options) => {
				options.suspicion.active = false;
				return options;
			}),
		setSuspicionRadius: (radius: string) =>
			update((options) => {
				const _radius = parseInt(radius);

				if (
					!isNaN(_radius) &&
					_radius >= Config.options.suspicion.radius.min &&
					_radius <= Config.options.suspicion.radius.max
				)
					options.suspicion.radius = _radius;

				return options;
			}),
		// ----------------------- set analyzer
		activeAnalyzerLayered: () =>
			update((options) => {
				options.analyzer.layered = true;
				return options;
			}),
		disableAnalyzerLayered: () =>
			update((options) => {
				options.analyzer.layered = false;
				return options;
			}),
		activeAnalyzerPercentage: () =>
			update((options) => {
				options.analyzer.percentage = true;
				return options;
			}),
		disableAnalyzerPercentage: () =>
			update((options) => {
				options.analyzer.percentage = false;
				return options;
			}),
		activeAnalyzerRounded: () =>
			update((options) => {
				options.analyzer.rounded = true;
				return options;
			}),
		disableAnalyzerRounded: () =>
			update((options) => {
				options.analyzer.rounded = false;
				return options;
			})
	};
}

const optionsStore = createOptionsStore();

export default optionsStore;
