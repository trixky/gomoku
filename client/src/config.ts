import { env } from '$env/dynamic/public';

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
				default: 2,
				min: 0,
				max: 10
			},
			pruning: {
				default: true
			}
		},
		poximity: {
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
			}
		}
	}
};
