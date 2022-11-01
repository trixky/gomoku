import { writable } from 'svelte/store';

function createTimeStore() {
	const { subscribe, set } = writable(0);

	return {
		subscribe,
		set,
		reset: () => set(0)
	};
}

const timeStore = createTimeStore();

export default timeStore;
