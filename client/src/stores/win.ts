import { writable } from 'svelte/store';
import type WinModel from '../models/win';

function generateWin(): WinModel {
	return <WinModel>{
		player: 0,
		methode: 'alignement'
	};
}

function createWinStore() {
	const { subscribe, set, update } = writable(<WinModel>generateWin());

	return {
		subscribe,
		set,
		reset: () => {
			set(generateWin());
		},
		confirmAlignement: () => {
			update((win) => {
				win.loophole = false;
				return win;
			});
		},
		setCapture: (player: 1 | 2) => {
			set(<WinModel>{
				player,
				methode: 'capture',
				loophole: false
			});
		},
		setAlignement: (player: 1 | 2, loophole: boolean) => {
			set(<WinModel>{
				player,
				methode: 'alignement',
				loophole
			});
		}
	};
}

const winStore = createWinStore();

export default winStore;
