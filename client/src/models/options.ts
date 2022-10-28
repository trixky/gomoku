export default interface Options {
	timeout: number; // ms
	depth: {
		max: number;
		min: number;
		pruning: boolean;
		reduction: boolean;
	};
	proximity: {
		radius: number;
		threshold: number;
		shape: number;
		show: boolean;
		evolution: boolean;
	};
}
