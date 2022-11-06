import { env } from '$env/dynamic/public';
import { SHAPES } from './models/algo_options';
const size = 19;

export default {
	environment: {
		api: {
			origin: env.PUBLIC_DOMAIN ?? 'localhost',
			port: env.PUBLIC_API_PORT
		}
	},
	goban: {
		width: size,
		height: size,
		piece_number: Math.pow(size, 2),
		cell_number: Math.pow(size - 1, 2)
	},
	options: {
		time_out: {
			step: 100,
			default: 500,
			min: 100, // ms
			max: 60000 // ms
		},
		ai: {
			default: 'threepio'
		},
		depth: {
			max: {
				default: 2,
				min: 1,
				max: 10
			},
			min: {
				default: 0,
				min: 0,
				max: 10
			},
			pruning: {
				default: false
			},
			pruning_percentage: {
				step: 10,
				default: 0,
				min: 0,
				max: 1000
			}
		},
		width: {
			max: {
				step: 10,
				default: 361,
				min: 1,
				max: 361
			},
			pruning: {
				default: false
			},
			pruning_percentage: {
				step: 10,
				default: 0,
				min: 0,
				max: 1000
			},
			multi_threading: {
				default: true
			}
		},
		proximity: {
			radius: {
				default: 2,
				min: 1,
				max: 9
			},
			threshold: {
				default: 2,
				min: 1,
				max: 36
			},
			show: {
				default: false
			},
			shape: {
				default: SHAPES.neighbour
			},
			evolution: {
				default: true
			},
			reduction: {
				default: false
			}
		},
		heuristics: {
			values: {
				min: 0,
				max: 10,
				default: 5
			},
			show: {
				default: false
			}
		},
		suspicion: {
			active: {
				default: false
			},
			radius: {
				min: 1,
				max: 19,
				default: 4
			}
		},
		analyzer: {
			layered: {
				default: true
			},
			percentage: {
				default: false
			},
			rounded: {
				default: true
			}
		}
	}
};
