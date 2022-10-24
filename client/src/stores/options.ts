import Config from '../config';
import { writable } from 'svelte/store';
import type OptionsModel from '../models/options';

function generateOptions(): OptionsModel {
	return <OptionsModel>{
		proximity: {
			radius: Config.options.poximity.radius.default,
			threshold: Config.options.poximity.threshold.default,
			show: Config.options.poximity.show.default
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
					_radius >= Config.options.poximity.radius.min &&
					_radius <= Config.options.poximity.radius.max
				)
					options.proximity.radius = _radius;
				return options;
			}),
		setThreshold: (threshold: string) =>
			update((options) => {
				const _threshold = parseInt(threshold);

				if (
					!isNaN(_threshold) &&
					_threshold >= Config.options.poximity.threshold.min &&
					_threshold <= Config.options.poximity.threshold.max
				) {
					options.proximity.threshold = _threshold;
				}
				return options;
			})
	};
}

const optionsStore = createOptionsStore();

export default optionsStore;
