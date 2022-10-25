package doubleThree

import "fmt"

func isLeftCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x, y-3, player) &&
		coordOpponent(goban, x, y-2, opponent) &&
		coordOpponent(goban, x, y-1, opponent) {
		return true
	}
	return false
}

func isRightCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x, y+3, player) &&
		coordOpponent(goban, x, y+2, opponent) &&
		coordOpponent(goban, x, y+1, opponent) {
		return true
	}
	return false
}

func isHorizontalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isLeftCapture(goban, x, y, player, opponent) ||
		isRightCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

func isUpCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y, player) &&
		coordOpponent(goban, x-2, y, opponent) &&
		coordOpponent(goban, x-1, y, opponent) {
		return true
	}
	return false
}

func isDownCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y, player) &&
		coordOpponent(goban, x+2, y, opponent) &&
		coordOpponent(goban, x+1, y, opponent) {
		return true
	}
	return false
}

func isVerticalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isUpCapture(goban, x, y, player, opponent) ||
		isDownCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

func isUpLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y-3, player) &&
		coordOpponent(goban, x-2, y-2, opponent) &&
		coordOpponent(goban, x-1, y-1, opponent) {
		return true
	}
	return false
}

func isDownLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y+3, player) &&
		coordOpponent(goban, x+2, y+2, opponent) &&
		coordOpponent(goban, x+1, y+1, opponent) {
		return true
	}
	return false
}

func isLeftDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isUpLeftDiagonalCapture(goban, x, y, player, opponent) ||
		isDownLeftDiagonalCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

func isDownRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x+3, y-3, player) &&
		coordOpponent(goban, x+2, y-2, opponent) &&
		coordOpponent(goban, x+1, y-1, opponent) {
		return true
	}
	return false
}

func isUpRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if coordPlayer(goban, x-3, y+3, player) &&
		coordOpponent(goban, x-2, y+2, opponent) &&
		coordOpponent(goban, x-1, y+1, opponent) {
		return true
	}
	return false
}

func isRightDiagonalCapture(goban Goban, x int, y int, player int, opponent int) bool {
	if isDownRightDiagonalCapture(goban, x, y, player, opponent) ||
		isUpRightDiagonalCapture(goban, x, y, player, opponent) {
		return true
	}
	return false
}

func isCapture(goban Goban, x int, y int, player int) bool {
	opponent := 254 // do before, for now here
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
