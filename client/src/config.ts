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
	board: {
		width: size,
		height: size,
		piece_number: Math.pow(size, 2),
		cell_number: Math.pow(size - 1, 2)
	},
	options: {
		timeout: {
			min: 100, // ms
			max: 60000 // ms
		},
		depth: {
			max: {
				default: 2,
				min: 0,
				max: 10
			},
			min: {
				default: 0,
				min: 0,
				max: 10
			},
			pruning: {
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
			no_update: {
				default: false
			}
		}
	}
};
