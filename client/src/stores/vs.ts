import { writable } from 'svelte/store';

enum Opponents {
	'ai',
	'human'
}

export const Modes = <Array<keyof typeof Opponents>>['ai', 'human'];

function createVsStore() {
	const { subscribe, set } = writable(<keyof typeof Opponents>Modes[0]);

	return {
		subscribe,
		set: (mode: keyof typeof Opponents) => {
			set(mode);
		},
		setAi: () => set(Modes[0]),
		setHuman: () => set(Modes[1])
	};
}

export const vsStore = createVsStore();
