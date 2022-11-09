export interface PlayerInfo {
	pieces: number;
	captures: number;
	alignement: boolean;
	win: boolean;
}

export interface PlayersInfo {
	player_1: PlayerInfo;
	player_2: PlayerInfo;
}
