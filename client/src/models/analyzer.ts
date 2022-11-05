export default interface Analyzer {
	layers: {
		cutted_by_max_width: number;
		cutted_by_time_out: number;
		pruned_in_depth: number;
		pruned_in_width: number;
		saved_by_min_depth: number;
		selected: number;
	}[];
}
