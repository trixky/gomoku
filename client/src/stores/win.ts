import { writable } from 'svelte/store';
import type WinModel from '../models/win';

function generateWin(): WinModel {
	return <WinModel>{
		player: 0,
		methode: 'alignement'
	};
}

const g_history: WinModel[] = [];

function createWinStore() {
	const { subscribe, set, update } = writable(<WinModel>generateWin());

	return {
		subscribe,
		set,
		reset: () => {
			g_history.splice(0, g_history.length);
			set(generateWin());
		},
		confirmAlignement: () => {
			update((win) => {
				g_history.push(<WinModel>{
					loophole: win.loophole,
					methode: win.methode,
					player: win.player
				});
				win.loophole = false;
				return win;
			});
		},
		setCapture: (player: 1 | 2) => {
			update((win) => {
				g_history.push(<WinModel>{
					loophole: win.loophole,
					methode: win.methode,
					player: win.player
				});

				return <WinModel>{
					player,
					methode: 'capture',
					loophole: false
				};
			});
		},
		setAlignement: (player: 1 | 2, loophole: boolean) => {
			update((win) => {
				g_history.push(<WinModel>{
					loophole: win.loophole,
					methode: win.methode,
					player: win.player
				});

				return <WinModel>{
					player,
					methode: 'alignement',
					loophole
				};
			});
		},
		undo: () => {
			const previous_win = g_history.pop();

			if (previous_win != undefined) {
				set(previous_win);
			}
		}
	};
}

const winStore = createWinStore();

export default winStore;
