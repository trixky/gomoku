const size = 19;

export default {
	board: {
		width: size,
		height: size,
		piece_number: Math.pow(size, 2),
		cell_number: Math.pow(size - 1, 2)
	}
};
