export default interface Options {
	timeout: number; // ms
	depth: {
		max: number;
		min: number;
		pruning: boolean;
		reduction: boolean;
	};
	width: {
		max: number;
		multi_threading: boolean;
		pruning: boolean;
	};
	proximity: {
		radius: number;
		threshold: number;
		shape: number;
		show: boolean;
		evolution: boolean;
	};
	heuristics: {
		potential_alignement: number;
		alignement: number;
		potential_capture: number;
		capture: number;
		random: number;
	};
}
