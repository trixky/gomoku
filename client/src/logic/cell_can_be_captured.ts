import type CellModel from '../models/cell';

export default function cellCanBeCaptured(x: number, y: number, cells: CellModel[][]): boolean {
	const player: 0 | 1 | 2 = cells[y][x].player;

	function is_in_goban(z: number) {
		return z >= 0 && z < 19;
	}

	if (player > 0) {
		const opponent = (player % 2) + 1;

		for (let i = -1; i < 2; i++)
			for (let j = -1; j < 2; j++) {
				if (i === 0 && j === 0) continue;

				const xp1i = x + 1 * i;
				if (!is_in_goban(xp1i)) continue;
				const yp1j = y + 1 * j;
				if (!is_in_goban(yp1j)) continue;

				if (cells[yp1j][xp1i].player === player) {
					const xm1i = x - 1 * i;
					if (!is_in_goban(xm1i)) continue;
					const ym1j = y - 1 * j;
					if (!is_in_goban(ym1j)) continue;
					const xp2i = x + 2 * i;
					if (!is_in_goban(xp2i)) continue;
					const yp2j = y + 2 * j;
					if (!is_in_goban(yp2j)) continue;

					if (
						(cells[ym1j][xm1i].player === opponent && cells[yp2j][xp2i].player === 0) ||
						(cells[ym1j][xm1i].player === 0 && cells[yp2j][xp2i].player === opponent)
					)
						return true;
				}
			}
	}

	return false;
}
