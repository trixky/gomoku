export default interface Response {
	goban: string;
	options: {
		time: number;
		position: {
			x: number;
			y: number;
		};
	};
}
