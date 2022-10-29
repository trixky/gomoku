import Config from '../config';
import type { AlgoOptions as AlgoOptionsModel } from '../models/algo_options';

function postNext(board: string, options: AlgoOptionsModel): Promise<string> {
	return new Promise<string>((resolve, reject) => {
		const url = `http://${Config.environment.api.origin}:${Config.environment.api.port}/next`;

		const data = {
			options,
			goban: board
		};

		fetch(url, {
			method: 'POST',
			body: JSON.stringify(data)
		})
			.then((response) => {
				response
					.text()
					.then((body) => {
						resolve(body);
					})
					.catch((err) => {
						reject(err);
					});
			})
			.catch((err) => {
				reject(err);
			});
	});
}

export default postNext;
