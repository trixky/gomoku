import type { PlayersInfo as PlayersInfoModel } from '../models/players_info';

export default function winByCapture(players_info: PlayersInfoModel): 0 | 1 | 2 {
	let winner: 0 | 1 | 2 = 0;

	if (players_info.player_1.captures >= 10) {
		winner = 1;
	} else if (players_info.player_2.captures >= 10) {
		winner = 2;
	}

	return winner;
}
