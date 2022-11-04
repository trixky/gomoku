package doubleThree

import m "github.com/trixky/gomoku/models"

type Goban [19][19]byte

// coordOnGoban evaluates if the coordinates are on the goban
func coordOnGoban(x int, y int) bool {
	if x < 0 || x > 18 || y < 0 || y > 18 {
		return false
	}
	return true
}

// coordOnGobanPos evaluates if the coordinates are on the goban
func coordOnGobanPos(pos m.Position) bool {
	if pos.X < 0 || pos.X > 18 || pos.Y < 0 || pos.Y > 18 {
		return false
	}
	return true
}

// coordPlayer evaluates if the value of the coordinates on the goban has player number
func coordPlayer(goban Goban, x int, y int, player uint8) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == byte(player) {
			return true
		}
	}
	return false
}

// coordPlayer evaluates if the value of the coordinates on the goban has player number
func coordPlayerPos(goban Goban, pos m.Position, player uint8) bool {
	if coordOnGobanPos(pos) == true {
		if goban[pos.X][pos.Y] == byte(player) {
			return true
		}
	}
	return false
}

// coordUnoccupied test if the value of the coordinates on the goban is zero (empty)
func coordUnoccupied(goban Goban, x int, y int) bool {
	if coordOnGoban(x, y) == true {
		if goban[x][y] == 0 {
			return true
		}
	}
	return false
}

// coordUnoccupied test if the value of the coordinates on the goban is zero (empty)
func coordUnoccupiedPos(goban Goban, pos m.Position) bool {
	if coordOnGobanPos(pos) == true {
		if goban[pos.X][pos.Y] == 0 {
			return true
		}
	}
	return false
}
