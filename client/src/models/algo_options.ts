export const SHAPES = {
	neighbour: 2,
	star: 1,
	square: 0
};

export interface AlgoOptions {
	time_out: number; // ms
	position: {
		x: number;
		y: number;
	};
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
