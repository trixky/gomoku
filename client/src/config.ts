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
