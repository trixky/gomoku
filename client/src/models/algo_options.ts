export const SHAPES = {
	square: 0,
	star: 1
};

export interface AlgoOptions {
	timeout: number; // ms
	position: {
		x: number;
		y: number;
	};
	depth: {
		max: number;
		min: number;
		pruning: boolean;
	};
	proximity: {
		radius: number;
		threshold: number;
		shape: number;
		evolution: boolean;
	};
}
