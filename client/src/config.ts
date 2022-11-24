import { env } from '$env/dynamic/public';
import { SHAPES } from './models/algo_options';
const size = 19;

export default {
	environment: {
		api: {
			url: env.PUBLIC_API_URL ?? 'http://localhost',
			port: env.PUBLIC_API_PORT
		}
	},
	goban: {
		width: size,
		height: size,
		piece_number: Math.pow(size, 2),
		cell_number: Math.pow(size - 1, 2)
	},
	animation: {
		fade: {
			slow: { duration: 250 },
			normal: { duration: 120 },
			fast: { duration: 75 }
		}
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
				default: 3,
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
				max: 400,
				orange: 280,
				red: 300
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
				max: 200,
				orange: 80,
				red: 100
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
				default: SHAPES.square
			},
			evolution: {
				default: true
			}
		},
		heuristics: {
			weights: {
				alignement: {
					min: 0,
					max: 10,
					default: 5
				},
				capture: {
					min: 0,
					max: 10,
					default: 2
				},
				random: {
					min: 0,
					max: 10,
					default: 0
				}
			},
			aggro: {
				min: 0,
				max: 300,
				default: 50,
				step: 10,
				orange: 90,
				red: 120
			},
			show: {
				default: false
			},
			depth_divisor: {
				default: 10,
				min: 0,
				max: 100,
				step: 1
			}
		},
		suggestion: {
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
