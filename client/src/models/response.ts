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
}
