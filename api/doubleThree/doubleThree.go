package doubleThree

import (
	m "github.com/trixky/gomoku/models"
)

// posArrayTranslation returns an array of positions that result of translations of 'pos' by (mulX, mulY). direction is for utility purpose
func posArrayTranslation(pos m.Position, mulX int, mulY int, direction int) []m.Position {
	translations := []m.Position{}
	for i := 1; i <=4; i++ {
		translation := pos
		translation.X += mulX * i * direction
		translation.Y += mulY * i * direction
		translation = append(translations, translation)
	}
	return translations
}

// checkEndThree returns boolean dep. if the position given would create a three w/ position in one of the ends
func checkEndThree(goban Goban, minus []m.Postion, plus []m.Postion, player int) {
	if coordUnoccupied(goban, plus[0]) {
		if coordPlayer(minus[0] player) {
			if coordPlayer(minus[1] player) && coordUnoccupied(goban, minus[2]) {
				return true // 011X0
			}
			if coordUnoccupied(goban, minus[1]) && coordPlayer(minus[2] player) && coordUnoccupied(goban, minus[3]) {
				return true // 0101X0
			}
		}
		if coordUnoccupied(goban, minus[0]) && coordPlayer(minus[1] player) && coordPlayer(minus[2] player) && coordUnoccupied(goban, minus[3]) {
			return true // 0110X0
		}
	}

	// just reverse positions array with the few lines up
	if coordUnoccupied(goban, minus[0]) {
		if coordPlayer(plus[0] player) {
			if coordPlayer(plus[1] player) && coordUnoccupied(goban, plus[2]) {
				return true // 0X110
			}
			if coordUnoccupied(goban, plus[1]) && coordPlayer(plus[2] player) && coordUnoccupied(goban, plus[4]) {
				return true // 0X1010
			}
		}
		if coordUnoccupied(goban, plus[0]) && coordPlayer(plus[1] player) && coordPlayer(plus[2] player) && coordUnoccupied(goban, plus[4]) {
			return true // 0X0110
		}
	}

	return false
}

// checMiddleThree returns boolean dep. if the position given would create a three w/ position in the middle
func checMiddleThree(goban Goban, minus []m.Postion, plus []m.Postion, player int) {
	if coordPlayer(plus[0] player) && coordUnoccupied(goban, plus[1]) {
		if coordPlayer(minus[0] player) && coordUnoccupied(goban, minus[1]) {
			return true // 01X10
		}
		if coordUnoccupied(goban, minus[0]) && coordPlayer(minus[1] player) && coordUnoccupied(goban, minus[2]) {
			return true // 010X10
		}
	}
	if coordUnoccupied(goban, plus[0]) && coordPlayer(plus[1] player) && coordUnoccupied(goban, plus[2]) {
		if coordPlayer(minus[0] player) && coordUnoccupied(goban, minus[1]) {
			return true // 01X010
		}
	}

	return false
}

// checkThree returns boolean dep. if the position given would create a free three w/ given direction
func checkThree(goban Goban, pos m.Position, player int, mulX int, mulY int) bool {
	minus := posArrayTranslation(pos, mulX, mulY, -1)
	plus := posArrayTranslation(pos, mulX, mulY, 1)

	endThree := checkEndThree(goban, pos, minus, plus, player)
	middleThree := checkMidddleThree(goban, pos, minus, plus, player)

	if (endThree || middleThree) {
		return true
	}
	return false
}

// checkDoubleThree returns boolean true if the position given would create a doubleThree, false otherwise
func checkDoubleThree(goban Goban, pos m.Position, player int) bool {
	hor := checkThree(goban, pos, player, 1, 0)
	vert := checkThree(goban, pos, player, 0, 1)
	lDiag := checkThree(goban, pos, player, 1, 1)
	rDiag := checkThree(goban, pos, player, 1, -1)

	if (hor && (vert || lDiag || rDiag)) ||
		(vert && (lDiag || rDiag)) ||
		(lDiag && rDiag) { // check at least two true
		return true
	}
	return false
}
