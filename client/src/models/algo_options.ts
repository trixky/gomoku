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
		pruning_percentage: number;
	};
	width: {
		max: number;
		pruning_percentage: number;
		multi_threading: boolean;
	};
	proximity: {
		radius: number;
		threshold: number;
		shape: number;
		evolution: boolean;
	};
	heuristics: {
		alignement: number;
		capture: number;
		random: number;
	};
	suspicion: {
		radius: number;
	};
}
