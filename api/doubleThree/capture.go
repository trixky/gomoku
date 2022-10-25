package doubleThree

import "fmt"

// isLeftCapture tests if current position would make a left capture
// x o o X
func isLeftCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x, y-3, player) &&
		coordPlayer(goban, x, y-2, opponent) &&
		coordPlayer(goban, x, y-1, opponent) {
		return true
	}
	return false
}

// isRightCapture tests if current position would make a right capture
// X o o x
func isRightCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x, y+3, player) &&
		coordPlayer(goban, x, y+2, opponent) &&
		coordPlayer(goban, x, y+1, opponent) {
		return true
	}
	return false
}

// isHorizontalCapture tests if current position would make a horizontal capture (see two func above)
func isHorizontalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isLeftCapture(goban, x, y, player, opponent) ||
		isRightCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

// isUpCapture tests if current position would make an upwards capture
// x
// o
// o
// X
func isUpCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y, player) &&
		coordPlayer(goban, x-2, y, opponent) &&
		coordPlayer(goban, x-1, y, opponent) {
		return true
	}
	return false
}

// isDownCapture tests if current position would make a downwards capture
// X
// o
// o
// x
func isDownCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y, player) &&
		coordPlayer(goban, x+2, y, opponent) &&
		coordPlayer(goban, x+1, y, opponent) {
		return true
	}
	return false
}

// isVerticalCapture tests if current position would make an vertical capture (see two func above)
func isVerticalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isUpCapture(goban, x, y, player, opponent) ||
		isDownCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

// isUpLeftDiagonalCapture tests if current position would make an upwards left diagonal capture
// x
//  o
//   o
//    X
func isUpLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y-3, player) &&
		coordPlayer(goban, x-2, y-2, opponent) &&
		coordPlayer(goban, x-1, y-1, opponent) {
		return true
	}
	return false
}

// isDownLeftDiagonalCapture tests if current position would make a downards left diagonal capture
// X
//  o
//   o
//    x
func isDownLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y+3, player) &&
		coordPlayer(goban, x+2, y+2, opponent) &&
		coordPlayer(goban, x+1, y+1, opponent) {
		return true
	}
	return false
}

// isDownLeftDiagonalCapture tests if current position would make a left diagonal capture (see two func above)
func isLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isUpLeftDiagonalCapture(goban, x, y, player, opponent) ||
		isDownLeftDiagonalCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

// isDownRightDiagonalCapture tests if current position would make a downards right diagonal capture
//    X
//   o
//  o
// x
func isDownRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y-3, player) &&
		coordPlayer(goban, x+2, y-2, opponent) &&
		coordPlayer(goban, x+1, y-1, opponent) {
		return true
	}
	return false
}

// isUpRightDiagonalCapture tests if current position would make an upwards right diagonal capture
//    x
//   o
//  o
// X
func isUpRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y+3, player) &&
		coordPlayer(goban, x-2, y+2, opponent) &&
		coordPlayer(goban, x-1, y+1, opponent) {
		return true
	}
	return false
}

// isRightDiagonalCapture tests if current position would make a right diagonal capture (see two func above)
func isRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isDownRightDiagonalCapture(goban, x, y, player, opponent) ||
		isUpRightDiagonalCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

// isCapture tests if current position would make a capture in any direction (see all func above)
func isCapture(goban Goban, x int, y int, player int) bool {
	opponent := 254
	if player == 254 {
		opponent++
	}
	if isHorizontalCapture(goban, x, y, player, opponent) {
		fmt.Println("hor false")
		return true
	} else if isVerticalCapture(goban, x, y, player, opponent) {
		return true
	} else if isLeftDiagonalCapture(goban, x, y, player, opponent) {
		return true
	} else if isRightDiagonalCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}
