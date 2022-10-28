import Config from '../config';
import { writable } from 'svelte/store';
import type OptionsModel from '../models/options';
import { SHAPES } from '../models/algo_options';

function generateOptions(): OptionsModel {
	return <OptionsModel>{
		proximity: {
			radius: Config.options.proximity.radius.default,
			threshold: Config.options.proximity.threshold.default,
			show: Config.options.proximity.show.default,
			shape: SHAPES.square
		}
	};
}

function createOptionsStore() {
	const { subscribe, set, update } = writable(generateOptions());

	return {
		subscribe,
		reset: () => set(generateOptions()),
		show: () =>
			update((options) => {
				options.proximity.show = true;
				return options;
			}),
		hide: () =>
			update((options) => {
				options.proximity.show = false;
				return options;
			}),
		setRadius: (radius: string) =>
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
		setThreshold: (threshold: string) =>
			update((options) => {
				const _threshold = parseInt(threshold);

				if (
					!isNaN(_threshold) &&
					_threshold >= Config.options.proximity.threshold.min &&
					_threshold <= Config.options.proximity.threshold.max
				) {
					options.proximity.threshold = _threshold;
				}
				return options;
			}),
		setShape: (shape: number) =>
			update((options) => {
				if (Object.values(SHAPES).includes(shape)) {
					options.proximity.shape = shape;
				}
				return options;
			})
	};
}

const optionsStore = createOptionsStore();

export default optionsStore;
