package doubleThree

import (
	m "github.com/trixky/gomoku/models"
)

type Goban [19][19]byte

// coordOnGoban test if the coordinates are on the goban
func coordOnGoban(pos m.Position) bool {
	if pox.X < 0 || pox.X > 18 || pos.Y < 0 || pos.Y > 18 {
		return false
	}
	return true
}

// coordPlayer test if the value of the coordinates on the goban has player number
func coordPlayer(goban Goban, pos m.Position player int) bool {
	if coordOnGoban(pos) == true {
		if goban[pos.X][pos.Y] == byte(player) {
			return true
		}
	}
	return false
}

// coordUnoccupied test if the value of the coordinates on the goban is zero (empty)
func coordUnoccupied(goban Goban, pos m.Position) bool {
	if coordOnGoban(pos) == true {
		if goban[pos.X][pos.Y] == 0 {
			return true
		}
	}
	return false
}
