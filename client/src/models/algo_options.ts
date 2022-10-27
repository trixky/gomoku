export const SHAPE_SQUARE = 0;
export const SHAPE_STAR = 1;

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
		no_update: boolean;
	};
}
