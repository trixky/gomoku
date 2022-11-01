import BoardStore from './goban';
import { derived } from 'svelte/store';

export default derived(BoardStore, ($BoardStore): string => {
	return $BoardStore.cells
		.map((cells) => cells.map((cell) => cell.player.toString()).join(''))
		.join('');
});
