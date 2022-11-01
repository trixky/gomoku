import { writable } from 'svelte/store';

function createLoadingStore() {
	const { subscribe, set } = writable(false);

	return {
		subscribe,
		switch: (on: boolean) => set(on)
	};
}

const loadingStore = createLoadingStore();

export default loadingStore;
