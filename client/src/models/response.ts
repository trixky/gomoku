import type Analyzer from './analyzer';
import type { PlayersInfo as PlayersInfoModel } from './players_info';

export default interface Response {
	goban: string;
	heuristic_goban: string;
	options: {
		time: number;
		position: {
			x: number;
			y: number;
		};
	};
	analyzer: Analyzer;
	players_info: PlayersInfoModel;
}
