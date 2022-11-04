export default interface Options {
	time_out: number; // ms
	depth: {
		max: number;
		min: number;
		pruning: boolean;
		pruning_percentage: number;
	};
	width: {
		max: number;
		min: number;
		multi_threading: boolean;
		pruning: boolean;
		pruning_percentage: number;
	};
	proximity: {
		radius: number;
		threshold: number;
		shape: number;
		show: boolean;
		evolution: boolean;
		reduction: boolean;
	};
	heuristics: {
		potential_alignement: number;
		alignement: number;
		potential_capture: number;
		capture: number;
		random: number;
		show: boolean;
	};
	suspicion: {
		active: boolean;
		radius: number;
	};
}
