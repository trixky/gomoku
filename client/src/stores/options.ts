import Config from '../config';
import { writable } from 'svelte/store';
import type OptionsModel from '../models/options';
import { SHAPES } from '../models/algo_options';

function generateOptions(): OptionsModel {
	return <OptionsModel>{
		depth: {
			max: Config.options.depth.max.default,
			min: Config.options.depth.min.default,
			pruning: Config.options.depth.pruning.default,
			reduction: Config.options.depth.reduction.default
		},
		proximity: {
			radius: Config.options.proximity.radius.default,
			threshold: Config.options.proximity.threshold.default,
			show: Config.options.proximity.show.default,
			evolution: Config.options.proximity.evolution.default,
			shape: SHAPES.square
		},
		heuristics: {
			potential_alignement: Config.options.heuristics.default,
			win_by_alignement: Config.options.heuristics.default,
			potential_capture: Config.options.heuristics.default,
			capture: Config.options.heuristics.default,
			random: Config.options.heuristics.default
		}
	};
}

function createOptionsStore() {
	const { subscribe, set, update } = writable(generateOptions());

	return {
		subscribe,
		reset: () => set(generateOptions()),
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
		setDepthReduction: (reduction: boolean) =>
			update((options) => {
				options.depth.reduction = reduction;
				return options;
			}),
		// ----------------------- set depth
		setHeuristicsPotentialAlignement: (potential_alignement: string) =>
			update((options) => {
				const _potential_alignement = parseInt(potential_alignement);

				if (!isNaN(_potential_alignement))
					options.heuristics.potential_alignement = _potential_alignement;
				return options;
			}),
		setHeuristicsWinByAlignement: (win_by_alignement: string) =>
			update((options) => {
				const _win_by_alignement = parseInt(win_by_alignement);

				if (!isNaN(_win_by_alignement)) options.heuristics.win_by_alignement = _win_by_alignement;
				return options;
			}),
		setHeuristicsPotentialCapture: (potential_capture: string) =>
			update((options) => {
				const _potential_capture = parseInt(potential_capture);

				if (!isNaN(_potential_capture)) options.heuristics.potential_capture = _potential_capture;
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
			})
	};
}

const optionsStore = createOptionsStore();

export default optionsStore;
