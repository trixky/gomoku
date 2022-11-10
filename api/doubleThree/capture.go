package doubleThree

import (
	m "github.com/trixky/gomoku/models"
)

// isLeftCapture tests if current position would make a left capture
// x o o X
func isLeftCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	horArray := []m.Position{}
	if coordPlayer(goban, x-3, y, player) &&
		coordPlayer(goban, x-2, y, opponent) &&
		coordPlayer(goban, x-1, y, opponent) {
		horArray = append(horArray, m.Position{X: uint8(x - 1), Y: uint8(y)}, m.Position{X: uint8(x - 2), Y: uint8(y)})
		return true, horArray
	}
	return false, horArray
}

// isRightCapture tests if current position would make a right capture
// X o o x
func isRightCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	horArray := []m.Position{}
	if coordPlayer(goban, x+3, y, player) &&
		coordPlayer(goban, x+2, y, opponent) &&
		coordPlayer(goban, x+1, y, opponent) {
		horArray = append(horArray, m.Position{X: uint8(x + 1), Y: uint8(y)}, m.Position{X: uint8(x + 2), Y: uint8(y)})
		return true, horArray
	}
	return false, horArray
}

// isHorizontalCapture tests if current position would make a horizontal capture (see two func above)
func isHorizontalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	truth, captureArray := false, []m.Position{}

	left, leftArray := isLeftCapture(goban, x, y, player, opponent)
	truth = truth || left
	captureArray = append(captureArray, leftArray...)

	right, rightArray := isRightCapture(goban, x, y, player, opponent)
	truth = truth || right
	captureArray = append(captureArray, rightArray...)

	return truth, captureArray
}

// isUpCapture tests if current position would make an upwards capture
// x
// o
// o
// X
func isUpCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	upArray := []m.Position{}
	if coordPlayer(goban, x, y-3, player) &&
		coordPlayer(goban, x, y-2, opponent) &&
		coordPlayer(goban, x, y-1, opponent) {
		upArray = append(upArray, m.Position{X: uint8(x), Y: uint8(y - 1)}, m.Position{X: uint8(x), Y: uint8(y - 2)})
		return true, upArray
	}
	return false, upArray
}

// isDownCapture tests if current position would make a downwards capture
// X
// o
// o
// x
func isDownCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	captArray := []m.Position{}
	if coordPlayer(goban, x, y+3, player) &&
		coordPlayer(goban, x, y+2, opponent) &&
		coordPlayer(goban, x, y+1, opponent) {
		captArray = append(captArray, m.Position{X: uint8(x), Y: uint8(y + 1)}, m.Position{X: uint8(x), Y: uint8(y + 2)})
		return true, captArray
	}
	return false, captArray
}

// isVerticalCapture tests if current position would make an vertical capture (see two func above)
func isVerticalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	truth, captureArray := false, []m.Position{}

	up, upArray := isUpCapture(goban, x, y, player, opponent)
	truth = truth || up
	if up {
		captureArray = append(captureArray, upArray...)
	}

	down, downArray := isDownCapture(goban, x, y, player, opponent)
	truth = truth || down
	if down {
		captureArray = append(captureArray, downArray...)
	}

	return truth, captureArray
}

// isUpLeftDiagonalCapture tests if current position would make an upwards left diagonal capture
// x
//  o
//   o
//    X
func isUpLeftDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	upLeft := []m.Position{}
	if coordPlayer(goban, x-3, y-3, player) &&
		coordPlayer(goban, x-2, y-2, opponent) &&
		coordPlayer(goban, x-1, y-1, opponent) {
		upLeft = append(upLeft, m.Position{X: uint8(x - 1), Y: uint8(y - 1)}, m.Position{X: uint8(x - 2), Y: uint8(y - 2)})
		return true, upLeft
	}
	return false, upLeft
}

// isDownRightDiagonalCapture tests if current position would make a downards left diagonal capture
// X
//  o
//   o
//    x
func isDownRightDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	upLeft := []m.Position{}
	if coordPlayer(goban, x+3, y+3, player) &&
		coordPlayer(goban, x+2, y+2, opponent) &&
		coordPlayer(goban, x+1, y+1, opponent) {
		upLeft = append(upLeft, m.Position{X: uint8(x + 1), Y: uint8(y + 1)}, m.Position{X: uint8(x + 2), Y: uint8(y + 2)})
		return true, upLeft
	}
	return false, upLeft
}

// isLeftDiagonalCapture tests if current position would make a left diagonal capture (see two func above)
func isLeftDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	truth, captureArray := false, []m.Position{}

	upLeft, upLeftArray := isUpLeftDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || upLeft
	if upLeft {
		captureArray = append(captureArray, upLeftArray...)
	}

	downRight, downRightArray := isDownRightDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || downRight
	if downRight {
		captureArray = append(captureArray, downRightArray...)
	}

	return truth, captureArray
}

// isDownLeftDiagonalCapture tests if current position would make a downards right diagonal capture
//    X
//   o
//  o
// x
func isDownLeftDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	downLeft := []m.Position{}
	if coordPlayer(goban, x-3, y+3, player) &&
		coordPlayer(goban, x-2, y+2, opponent) &&
		coordPlayer(goban, x-1, y+1, opponent) {
		downLeft = append(downLeft, m.Position{X: uint8(x - 1), Y: uint8(y + 1)}, m.Position{X: uint8(x - 2), Y: uint8(y + 2)})
		return true, downLeft
	}
	return false, downLeft
}

// isUpRightDiagonalCapture tests if current position would make an upwards right diagonal capture
//    x
//   o
//  o
// X
func isUpRightDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	upRight := []m.Position{}
	if coordPlayer(goban, x+3, y-3, player) &&
		coordPlayer(goban, x+2, y-2, opponent) &&
		coordPlayer(goban, x+1, y-1, opponent) {
		upRight = append(upRight, m.Position{X: uint8(x + 1), Y: uint8(y - 1)}, m.Position{X: uint8(x + 2), Y: uint8(y - 2)})
		return true, upRight
	}
	return false, upRight
}

// isRightDiagonalCapture tests if current position would make a right diagonal capture (see two func above)
func isRightDiagonalCapture(goban m.Goban, x int, y int, player uint8, opponent uint8) (bool, []m.Position) {
	truth, captureArray := false, []m.Position{}

	downLeft, downLeftArray := isDownLeftDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || downLeft
	if downLeft {
		captureArray = append(captureArray, downLeftArray...)
	}

	upRight, upRightArray := isUpRightDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || upRight
	if upRight {
		captureArray = append(captureArray, upRightArray...)
	}

	return truth, captureArray
}

// isCapture tests if current position would make a capture in any direction (see all func above)
func isCapture(goban m.Goban, x int, y int, player uint8) (bool, []m.Position) {
	var opponent uint8 = 254
	if player == 254 {
		opponent++
	}

	truth, captureArray := false, []m.Position{}

	hor, horArray := isHorizontalCapture(goban, x, y, player, opponent)
	truth = truth || hor
	if hor {
		captureArray = append(captureArray, horArray...)
	}

	vert, vertArray := isVerticalCapture(goban, x, y, player, opponent)
	truth = truth || vert
	if vert {
		captureArray = append(captureArray, vertArray...)
	}

	left, leftArray := isLeftDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || left

	if left {
		captureArray = append(captureArray, leftArray...)
	}

	right, rightArray := isRightDiagonalCapture(goban, x, y, player, opponent)
	truth = truth || right
	if right {
		captureArray = append(captureArray, rightArray...)
	}

	return truth, captureArray
}
